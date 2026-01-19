package ghost
 import(
	"fmt"
	"os"
 )
/*Quand un fichier est supprimé, il est déplacé dans le dossier caché ./ghost et ses métadonnés sont stockés dans un json */
func Delete(file string) error{
		// récuperer Les métadonnes du fichier
	info,err:= os.Stat(file)
	filename:= info.Name()

	ghostDir := "./.ghost/"
	dest := ghostDir + filename +".deleted" //la destination ou le fichier va etre deplace avec son nouveau nom

	err = os.Rename(file,dest) //Déplacer le fichier et le renommer
	if err != nil {fmt.Println(err)}
	return err
}