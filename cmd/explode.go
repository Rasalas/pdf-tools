/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/spf13/cobra"
)

// Funktion zum Aufteilen der PDF
func splitPDF(inputFile, destPattern string, start, end, pages int) error {
	ctx, err := api.ReadContextFile(inputFile)
	if err != nil {
		return fmt.Errorf("Fehler beim Lesen der PDF: %v", err)
	}

	// Bestimme die Endseite, wenn "*" angegeben ist
	if end == -1 || end > ctx.PageCount {
		end = ctx.PageCount
	}

	counter := 0
	for i := start + 1; i <= end; i += pages {
		counter++
		to := i + pages - 1
		if to > end {
			to = end
		}

		destFile := strings.Replace(destPattern, "#n#", strconv.Itoa(counter), -1)
		if !strings.HasSuffix(destFile, ".pdf") {
			destFile += ".pdf"
		}

		// Stelle sicher, dass das Zielverzeichnis existiert
		destDir := filepath.Dir(destFile)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("Fehler beim Erstellen des Zielverzeichnisses: %v", err)
		}

		pageRange := []string{fmt.Sprintf("%d-%d", i, to)}
		err = api.TrimFile(inputFile, destFile, pageRange, nil)
		if err != nil {
			return fmt.Errorf("Fehler beim Extrahieren der Seiten %d-%d: %v", i, to, err)
		}
		fmt.Printf("Seiten %d-%d gespeichert als %s\n", i, to, destFile)
	}

	return nil
}

// explodeCmd represents the explode command
var explodeCmd = &cobra.Command{
	Use:   "explode [inputfile]",
	Short: "Teile eine PDF in mehrere Dateien auf",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]

		destPattern, _ := cmd.Flags().GetString("dest")
		start, _ := cmd.Flags().GetInt("start")
		endStr, _ := cmd.Flags().GetString("end")
		pages, _ := cmd.Flags().GetInt("pages")

		// Standardmuster für Dateinamen festlegen, falls nicht angegeben
		if destPattern == "" {
			originalName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
			destPattern = originalName + "_shard-{n}.pdf"
		}

		// Verarbeite den "end"-Parameter
		var end int
		if endStr == "*" {
			end = -1 // -1 wird in splitPDF als "alle Seiten" interpretiert
		} else {
			var err error
			end, err = strconv.Atoi(endStr)
			if err != nil {
				log.Fatalf("Ungültiger Wert für --end: %v\n", err)
			}
		}

		// Split PDF
		err := splitPDF(inputFile, destPattern, start, end, pages)
		if err != nil {
			log.Fatalf("Fehler beim Aufteilen der PDF: %v\n", err)
		}
		fmt.Println("PDF wurde erfolgreich aufgeteilt!")
	},
}

func init() {
	rootCmd.AddCommand(explodeCmd)

	// Flags für das Subcommand "explode"
	explodeCmd.Flags().String("dest", "", "Das Muster für den Dateinamen der resultierenden PDFs (z.B. result_{n}.pdf)")
	explodeCmd.Flags().Int("start", 0, "Die Startseite (Index beginnt bei 0)")
	explodeCmd.Flags().String("end", "*", "Die Endseite (verwende * für alle Seiten)")
	explodeCmd.Flags().Int("pages", 1, "Anzahl der Seiten pro resultierender PDF")
}
