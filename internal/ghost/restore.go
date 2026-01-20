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
			fmt.Println("Erreur de lecture Json: ", err)
		}

		// parcourir le tableau et restaurer le fichier correspondant
		fmt.Printf("\n-> Réstauration du fichier !\n",)
		for index, value := range InfoTab {
			if index+1 == id{ //si le fichier existe
				src := "./.ghost/"+value.FileName+".deleted" //source du fichier supprimé
				dest := value.FilePath //detination
				err = os.Rename(src,dest)
				if err != nil {fmt.Println("Erreur de réstauration du fichier: ",err)}else{fmt.Printf("Fichier réstauré !\n")}

				// Mettre à jour le fichier json en supprimant l'info du fichier par son index
				InfoTab = append(InfoTab[:index],InfoTab[index+1:]...) //fusionne la partie inférieure et supérieur à l'index
				// Convertir le tableau en json
				jsonInfo, err := json.MarshalIndent(InfoTab,""," ")
				if err != nil {fmt.Println("Erreur de conversion JSON: ",err)}
				// Ecrire le json final dans le fichier metadata.json
				err = os.WriteFile("metadata.json",jsonInfo,0644)
				if err != nil {fmt.Println("Erreur d'écriture Json: ",err)}

			}
		}
	} else {
		fmt.Printf("\n-> Aucun fichier supprimé ! :\n")
	}
}