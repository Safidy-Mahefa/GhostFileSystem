package ghost

import(
	"fmt"
	"encoding/json"
	"os"
)
func List(){
	var InfoTab[] Info // le tableau qui contient l'info de chaque fichier

	jsonData,err:= os.ReadFile("metadata.json")
	if(len(jsonData) != 0){ //si je fichier json n'est pas vide
		fmt.Println("JsonData:", jsonData)
		err = json.Unmarshal([]byte(jsonData),&InfoTab) // pusher l'info json dans le tableau
		if err != nil {fmt.Println("Erreur de lecture Json: ",err)}

		// parcourir le tableau et afficher les fichiers supprimés
		fmt.Printf("\n-> Liste des fichiers supprimés :\n")
		for index,value:= range InfoTab{
			fmt.Println(index+1,":",value.FileName)
		}
	}else{
		fmt.Printf("\n-> Aucun fichier supprimé ! :\n")
	}
}