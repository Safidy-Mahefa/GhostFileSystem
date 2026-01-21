package ghost

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clean() {
	entries , err := os.ReadDir("./.ghost")
	if err != nil {ShowError(fmt.Sprintf("Erreur de lecture du dossier ghost: %v", err))}

	// Parcourir les elements du dossier et supprimer
	for _,entry := range entries{
		path := filepath.Join("./.ghost", entry.Name()) //pour avoir le chemin vers le fichier
		if path == "" || path == "."{
			ShowError("Chemin dangereux !")
			break
		}
		// Supprimer le fichier
		err = os.RemoveAll(path)
		if err != nil {ShowError(fmt.Sprintf("Erreur de suppression: %v", err))}

		// Mettre à jour le json
		err = os.WriteFile("metadata.json",[]byte{},0644)
		if err != nil {ShowError(fmt.Sprintf("Erreur lors du màj du json: %v", err))}
	}
}