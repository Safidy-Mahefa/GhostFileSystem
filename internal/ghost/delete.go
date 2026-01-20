package ghost

import (
	"encoding/json"
	"fmt"
	"os"
)
 type Info struct{ //une structure d'information pour chaque fichier
 	FileName string `json: "filename"`
	FileTime string `json: "fileTime"`
	FileSize int64 	`json: "fileSize"`
	FilePath string `json: "filePath"`
 }
/*Quand un fichier est supprimé, il est déplacé dans le dossier caché ./ghost et ses métadonnés sont stockés dans un json */
func Delete(file string) error{
	var InfoTab[] Info // le tableau qui contient l'info de chaque fichier

	// récuperer Les métadonnes du fichier,  on considère que file est le chemin originale du fichier....
	info,err:= os.Stat(file)
	var currentFileInfo Info //les infos du fichier courant
	currentFileInfo.FileName =  info.Name()
	currentFileInfo.FileTime =  info.ModTime().Format("02-01-2006 15:04") // .Format(exemple)
	currentFileInfo.FileSize =  info.Size()
	currentFileInfo.FilePath = file //le chemin original

		// Lire le fichier json
	jsonData,err:= os.ReadFile("metadata.json")
	if(len(jsonData) != 0){ //si je fichier json n'est pas vide
		fmt.Println("JsonData:", jsonData)
		err = json.Unmarshal([]byte(jsonData),&InfoTab) // pusher l'info json dans le tableau
		if err != nil {fmt.Println("Erreur de lecture Json: ",err)}
	}

	// Pusher l'info du fichier courant dans le tableau:
	InfoTab = append(InfoTab, currentFileInfo)

	// Convertir le tableau en json
	jsonInfo, err := json.MarshalIndent(InfoTab,""," ")
	if err != nil {fmt.Println("Erreur de conversion Json: ",err)}
	cleanJson := string(jsonInfo) //le json final
	fmt.Println("Json:", cleanJson)

	// Ecrire le json final dans le fichier metadata.json
	err = os.WriteFile("metadata.json",jsonInfo,0644)
	if err != nil {fmt.Println("Erreur d'écriture Json: ",err)}




	ghostDir := "./.ghost/"
	dest := ghostDir + currentFileInfo.FileName +".deleted" //la destination ou le fichier va etre deplace avec son nouveau nom

	err = os.Rename(file,dest) //Déplacer le fichier et le renommer
	if err != nil {fmt.Println("Erreur de déplacement:",err)}

	return err
}