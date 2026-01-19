package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"GSF/internal/ghost"
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
			ghost.Delete(file)
		}
	}
}