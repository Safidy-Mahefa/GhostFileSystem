package main

import (
	"fmt"
	"os"
	// "GSF/internal/ghost"
	// "GSF/internal/metadata"
	// "GSF/internal/utils"
)

func main() {
	boucle:= true
	commande := ""
	for boucle{ //boucle pour faire tourner le CLI
		fmt.Printf("-> ")
		v , err := fmt.Scanf("%s",&commande) //entréé utilisateur
		fmt.Println(v)
		fmt.Printf("\n")

		if err != nil { fmt.Println("Erreur :", err)}

		if commande == "q"{
			boucle = false
		}else if commande == "r"{
			os.Remove("file.txt")
		}
	}
}