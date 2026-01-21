package ghost

import (
	"encoding/json"
	"fmt"
	"os"
)

func Restore(id int) {
	var InfoTab []Info // le tableau qui contient l'info de chaque fichier

	jsonData, err := os.ReadFile("metadata.json")
	if len(jsonData) != 0 { //si je fichier json n'est pas vide
		err = json.Unmarshal([]byte(jsonData), &InfoTab) // pusher les infos json dans le tableau
		if err != nil {
			ShowError(fmt.Sprintf("Erreur de lecture Json: %v", err))
		}

		// parcourir le tableau et restaurer le fichier correspondant
		ShowInfo("Attempting to resurrect file from shadow realm...")
		for index, value := range InfoTab {
			if index+1 == id{ //si le fichier existe
				src := "./.ghost/"+value.FileName+".deleted" //source du fichier supprimé
				dest := value.FilePath //detination
				err = os.Rename(src,dest)
				if err != nil {
					ShowError(fmt.Sprintf("Erreur de réstauration du fichier: %v", err))
				}else{
					ShowSuccess(fmt.Sprintf("File '%s' successfully resurrected!", value.FileName))
				}

				// Mettre à jour le fichier json en supprimant l'info du fichier par son index
				InfoTab = append(InfoTab[:index],InfoTab[index+1:]...) //fusionne la partie inférieure et supérieur à l'index
				// Convertir le tableau en json
				jsonInfo, err := json.MarshalIndent(InfoTab,""," ")
				if err != nil {ShowError(fmt.Sprintf("Erreur de conversion JSON: %v", err))}
				// Ecrire le json final dans le fichier metadata.json
				err = os.WriteFile("metadata.json",jsonInfo,0644)
				if err != nil {ShowError(fmt.Sprintf("Erreur d'écriture Json: %v", err))}
				
				return
			}
		}
		ShowError(fmt.Sprintf("File with ID %d not found in shadow realm", id))
	} else {
		ShowWarning("Shadow realm is empty. No files to restore.")
	}
}