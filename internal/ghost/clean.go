package ghost

import (
	"fmt"
	"os"
	"path/filepath"
)

func Clean() {
	entries , err := os.ReadDir("./.ghost")
	if err != nil {fmt.Println("Erreur de lecture du dossier ghost: ",err)}
	fmt.Println(entries)

	// Parcourir les elements du dossier et supprimer
	for _,entry := range entries{
		path := filepath.Join("./.ghost", entry.Name()) //pour avoir le chemin vers le fichier
		if path == "" || path == "."{
			fmt.Println("Chemin dangereux !")
			break
		}
		// Supprimer le fichier
		err = os.RemoveAll(path)
		if err != nil {fmt.Println("Erreur de suppression: ",err)}

		// Mettre à jour le json
		err = os.WriteFile("metadata.json",[]byte{},0644)
		if err != nil {fmt.Println("Erreur lors du màj du json: ",err)}
	}
}