package main

import (
	"GSF/internal/ghost"
	"bufio" //entréé utilisateur safe
	"fmt"
	"os"
	"strconv" //pour convertir un string en int,...
	"strings" //manipuler du string
	// "GSF/internal/metadata"
	// "GSF/internal/utils"
)

func main() {
	boucle:= true
	reader := bufio.NewReader(os.Stdin) //pour lire l'entrée utilisateur à pusieurs espaces
	var parts[]string // Le tableau qui contient les parts de la commande

	for boucle{ //boucle pour faire tourner le CLI
		// Input utilisateur
		fmt.Printf("-> ")
		commande , err := reader.ReadString('\n') //entréé utilisateur
		fmt.Printf("\n")

		parts = strings.Split(commande," ")

		if err != nil { fmt.Println("Erreur :", err)}

		// Gestion des commandes
		if commande == "q"{
			boucle = false
		}else if strings.Contains(commande,"remove") && parts[0] == "remove" { //si la commande contient "remove" et commence par remove,
			file:=strings.TrimSpace(parts[1]) //nettoie le nom du fichier
			fmt.Println(file)
			file = strings.Replace(file,`\`,"/",-1) // pour avoir un chemin valide et compatible... -1 pour remplacer tout les occurences
			ghost.Delete(file)
		}else if strings.Contains(commande,"list"){ //list
			ghost.List()
		}else if strings.Contains(commande,"restore") && parts[0] == "restore"{
			data := strings.TrimSpace(parts[1]) //nettoyer l'input utilisateur
			id,err := strconv.Atoi(data)
			if err != nil {fmt.Println("Erreur de conversion ASCII to int: ",err)}
			ghost.Restore(id)
		}else if strings.Contains(commande,"clean"){ //clean
			ghost.Clean()
		}

	}
}