# Release-Prozess

Um einen neuen Release für pdf-tools zu erstellen, folgen Sie diesen Schritten:

1. Stellen Sie sicher, dass alle gewünschten Änderungen in den main-Branch gemerged wurden.

2. Aktualisieren Sie die Versionsnummer in der README.md und anderen relevanten Dateien.

3. Committen Sie diese Änderungen:
   ````
   git add .
   git commit -m "Bump version to X.Y.Z"
   ````

4. Erstellen Sie einen neuen Git-Tag mit der Versionsnummer:
   ````
   git tag -a vX.Y.Z -m "Release version X.Y.Z"
   ````

5. Pushen Sie den Commit und den Tag zum GitHub-Repository:
   ````
   git push origin main
   git push origin vX.Y.Z
   ````

6. Der GitHub Actions Workflow wird automatisch gestartet, sobald der Tag gepusht wird. Er wird GoReleaser ausführen, um die Binärdateien zu erstellen und einen neuen Release auf GitHub zu veröffentlichen.

7. Überprüfen Sie den GitHub Actions Workflow und den erstellten Release auf der GitHub-Releases-Seite.

8. Aktualisieren Sie bei Bedarf die Release-Notizen auf der GitHub-Releases-Seite mit zusätzlichen Informationen oder Änderungen.

Hinweis: Stellen Sie sicher, dass Sie die erforderlichen Berechtigungen haben, um Releases zu erstellen und zu veröffentlichen.