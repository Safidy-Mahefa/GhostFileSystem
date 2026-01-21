package ghost

import(
	"encoding/json"
	"os"
)
func List(){
	var InfoTab[] Info // le tableau qui contient l'info de chaque fichier

	jsonData,err:= os.ReadFile("metadata.json")
	if(len(jsonData) != 0){ //si je fichier json n'est pas vide
		err = json.Unmarshal([]byte(jsonData),&InfoTab) // pusher l'info json dans le tableau
		if err != nil {ShowError("Erreur de lecture Json: " + err.Error())}

		// parcourir le tableau et afficher les fichiers supprimés
		ShowTable(InfoTab)
	}else{
		ShowWarning("Shadow realm is empty. No files to display.")
	}
}