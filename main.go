package main

import (
	"GSF/internal/ghost"
	"bufio" //entréé utilisateur safe
	"encoding/json"
	"fmt"
	"os"
	"strconv" //pour convertir un string en int,...
	"strings" //manipuler du string
	// "GSF/internal/metadata"
	// "GSF/internal/utils"
)

func main() {
	// Afficher la bannière et le menu au démarrage
	ghost.ShowBanner()
	ghost.ShowMenu()
	
	boucle:= true
	reader := bufio.NewReader(os.Stdin) //pour lire l'entrée utilisateur à pusieurs espaces
	var parts[]string // Le tableau qui contient les parts de la commande

	for boucle{ //boucle pour faire tourner le CLI
		// Calculer les stats pour le prompt
		fileCount, totalSize := getGhostStats()
		
		// Input utilisateur
		fmt.Print(ghost.GetPrompt(fileCount, totalSize))
		commande , err := reader.ReadString('\n') //entréé utilisateur
		fmt.Printf("\n")

		parts = strings.Split(commande," ")

		if err != nil { ghost.ShowError(fmt.Sprintf("Erreur : %v", err))}

		// Gestion des commandes
		if commande == "q" || strings.TrimSpace(commande) == "exit"{
			ghost.ShowExitMessage()
			boucle = false
		}else if strings.Contains(commande,"remove") && parts[0] == "remove" { //si la commande contient "remove" et commence par remove,
			file:=strings.TrimSpace(parts[1]) //nettoie le nom du fichier
			file = strings.Replace(file,`\`,"/",-1) // pour avoir un chemin valide et compatible... -1 pour remplacer tout les occurences
			ghost.ShowProgress("Encrypting", 10)
			ghost.ShowProgress("Moving to shadow realm", 10)
			err := ghost.Delete(file)
			if err != nil {
				ghost.ShowError(fmt.Sprintf("Failed to delete file: %v", err))
			} else {
				ghost.ShowSuccess("File successfully erased from existence")
			}
		}else if strings.Contains(commande,"list"){ //list
			ghost.List()
		}else if strings.Contains(commande,"restore") && parts[0] == "restore"{
			data := strings.TrimSpace(parts[1]) //nettoyer l'input utilisateur
			id,err := strconv.Atoi(data)
			if err != nil {ghost.ShowError(fmt.Sprintf("Erreur de conversion ASCII to int: %v", err))}
			ghost.Restore(id)
		}else if strings.Contains(commande,"clean"){ //clean
			if ghost.ShowCleanWarning() {
				ghost.ShowProgress("Purging shadow realm", 10)
				ghost.Clean()
				ghost.ShowSuccess("All shadows have been purged permanently")
			} else {
				ghost.ShowInfo("Operation aborted by operator")
			}
		}else if strings.Contains(commande,"help"){
			ghost.ShowMenu()
		}else if strings.TrimSpace(commande) != "" {
			ghost.ShowError("Unknown command. Type 'help' for available commands")
		}

	}
}

// getGhostStats calcule le nombre de fichiers et la taille totale
func getGhostStats() (int, int64) {
	var InfoTab []ghost.Info
	var totalSize int64 = 0
	
	jsonData, err := os.ReadFile("metadata.json")
	if err == nil && len(jsonData) != 0 {
		json.Unmarshal([]byte(jsonData), &InfoTab)
		for _, info := range InfoTab {
			totalSize += info.FileSize
		}
	}
	
	return len(InfoTab), totalSize
}