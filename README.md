# PDF-Tools

PDF-Tools ist eine Kommandozeilenanwendung, die entwickelt wurde, um große PDF-Dateien in kleinere Teile aufzuteilen. Das Hauptwerkzeug ist der "explode" Befehl, der eine einzelne PDF-Datei in mehrere kleinere PDF-Dateien aufteilt.

## Installation

Um PDF-Tools zu installieren, stellen Sie sicher, dass Go auf Ihrem System installiert ist. Dann führen Sie den folgenden Befehl aus:

```bash
go install github.com/rasalas/pdf-tools@latest
```

## Verwendung

Der Hauptbefehl ist `explode`. Hier ist die grundlegende Syntax:

```bash
pdf-tools explode [inputfile] [flags]
```

### Flags

- `--dest`: Das Muster für den Dateinamen der resultierenden PDFs (z.B. result_{n}.pdf)
- `--start`: Die Startseite (Index beginnt bei 0)
- `--end`: Die Endseite (verwende * für alle Seiten)
- `--pages`: Anzahl der Seiten pro resultierender PDF

### Beispiele

1. Teile eine PDF in einzelne Seiten auf:
   ```
   pdf-tools explode input.pdf
   ```

2. Teile eine PDF in 5-Seiten-Abschnitte auf:
   ```
   pdf-tools explode input.pdf --pages 5
   ```

3. Teile eine PDF von Seite 10 bis zum Ende in 3-Seiten-Abschnitte auf:
   ```
   pdf-tools explode input.pdf --start 9 --end * --pages 3
   ```

4. Teile eine PDF mit benutzerdefinierten Ausgabedateinamen auf:
   ```
   pdf-tools explode input.pdf --dest "output_{n}.pdf"
   ```

## Hinweise

- Die Seitennummerierung beginnt bei 0, nicht bei 1.
- Verwenden Sie `*` als `--end`-Wert, um bis zum Ende der PDF zu gehen.
- Das Ausgabemuster kann `#n#` enthalten, das durch die Teilungsnummer ersetzt wird.
