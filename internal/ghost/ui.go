package ghost

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Codes couleurs ANSI
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorGray   = "\033[90m"
	
	// Styles
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
)

// ShowBanner affiche la bannière ASCII au démarrage
func ShowBanner() {
	banner := `
   ▄████  ██░ ██  ▒█████    ██████ ▄▄▄█████▓
  ██▒ ▀█▒▓██░ ██▒▒██▒  ██▒▒██    ▒ ▓  ██▒ ▓▒
 ▒██░▄▄▄░▒██▀▀██░▒██░  ██▒░ ▓██▄   ▒ ▓██░ ▒░
 ░▓█  ██▓░▓█ ░██ ▒██   ██░  ▒   ██▒░ ▓██▓ ░ 
 ░▒▓███▀▒░▓█▒░██▓░ ████▓▒░▒██████▒▒  ▒██▒ ░ 
  ░▒   ▒  ▒ ░░▒░▒░ ▒░▒░▒░ ▒ ▒▓▒ ▒ ░  ▒ ░░   
   ░   ░  ▒ ░▒░ ░  ░ ▒ ▒░ ░ ░▒  ░ ░    ░    
 ░ ░   ░  ░  ░░ ░░ ░ ░ ▒  ░  ░  ░    ░      
       ░  ░  ░  ░    ░ ░        ░           
`
	fmt.Print(ColorCyan + Bold + banner + ColorReset)
	fmt.Println(ColorGray + "        Ghost Secure Files v1.0 | [SYSTEM ACTIVE]" + ColorReset)
	fmt.Println()
	
	// Easter egg aléatoire
	showRandomQuote()
	fmt.Println()
}

// showRandomQuote affiche une citation aléatoire
func showRandomQuote() {
	quotes := []string{
		"The files you delete... they're watching...",
		"In the shadow realm, nothing is truly gone.",
		"Every deletion leaves a trace in the void.",
		"Ghost protocol initiated. Welcome back, operator.",
		"Files don't die. They fade into darkness.",
		"The digital graveyard awaits your command.",
	}
	
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println(ColorGray + Italic + "   > " + quote + ColorReset)
}

// ShowMenu affiche le menu d'aide
func ShowMenu() {
	menu := `
╔═══════════════════════ GHOST COMMANDS ═══════════════════════╗
║                                                              ║
║  remove <file>    → Send file to shadow realm                ║
║  list             → Display archived entities                ║
║  restore <id>     → Resurrect file from ghost                ║
║  clean            → Purge all shadows [PERMANENT]            ║
║  help             → Show this transmission                   ║
║  q / exit         → Disconnect from system                   ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`
	fmt.Print(ColorCyan + menu + ColorReset)
}

// ShowSuccess affiche un message de succès
func ShowSuccess(message string) {
	fmt.Println(ColorGreen + Bold + "[✓]" + ColorReset + ColorGreen + " OPERATION COMPLETE → " + message + ColorReset)
}

// ShowError affiche un message d'erreur
func ShowError(message string) {
	fmt.Println(ColorRed + Bold + "[✗]" + ColorReset + ColorRed + " ACCESS DENIED → " + message + ColorReset)
}

// ShowWarning affiche un avertissement
func ShowWarning(message string) {
	fmt.Println(ColorYellow + Bold + "[!]" + ColorReset + ColorYellow + " CAUTION → " + message + ColorReset)
}

// ShowInfo affiche une information
func ShowInfo(message string) {
	fmt.Println(ColorCyan + Bold + "[i]" + ColorReset + ColorCyan + " INFO → " + message + ColorReset)
}

// ShowProgress simule une barre de progression
func ShowProgress(operation string, steps int) {
	bar := []string{"▓", "▓", "▓", "▓", "▓", "▓", "▓", "▓", "▓", "▓"}
	
	for i := 0; i <= steps; i++ {
		percentage := (i * 100) / steps
		filled := (i * 10) / steps
		
		progress := ""
		for j := 0; j < 10; j++ {
			if j < filled {
				progress += ColorGreen + bar[j] + ColorReset
			} else {
				progress += ColorGray + "░" + ColorReset
			}
		}
		
		fmt.Printf("\r[%s] %s %d%%", progress, operation, percentage)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}

// ShowTable affiche une table formatée des fichiers
func ShowTable(infoTab []Info) {
	if len(infoTab) == 0 {
		fmt.Println()
		ShowWarning("Shadow realm is empty. No files to display.")
		fmt.Println()
		return
	}
	
	fmt.Println()
	fmt.Println(ColorCyan + Bold +  "╔══════════════════════════════════════════════════════════════════════╗" + ColorReset)
	fmt.Println(ColorCyan + Bold + "║               SHADOW ARCHIVE — DELETED FILES                         ║" + ColorReset)
	fmt.Println(ColorCyan + "╠════╦═══════════════════════════╦═══════════╦═════════════════════════╣" + ColorReset)
	fmt.Println(ColorCyan + "║ ID ║ FILENAME                  ║ SIZE      ║ TIMESTAMP               ║" + ColorReset)
	fmt.Println(ColorCyan + "╠════╬═══════════════════════════╬═══════════╬═════════════════════════╣" + ColorReset)
	
	for index, value := range infoTab {
		// Formater la taille
		size := formatSize(value.FileSize)
		
		// Tronquer le nom si trop long
		filename := value.FileName
		if len(filename) > 25 {
			filename = filename[:22] + "..."
		}
		
		// Afficher la ligne
		fmt.Printf(ColorWhite+"║ "+ColorYellow+Bold+"%-2d"+ColorReset+ColorWhite+" ║ %-25s ║ %-9s ║ %-23s ║"+ColorReset+"\n",
			index+1,
			filename,
			size,
			value.FileTime,
		)
	}
	
	fmt.Println(ColorCyan + "╚════╩═══════════════════════════╩═══════════╩═════════════════════════╝" + ColorReset)
	fmt.Println()
}

// formatSize convertit la taille en format lisible
func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// ShowStats affiche les statistiques
func ShowStats(fileCount int, totalSize int64) {
	size := formatSize(totalSize)
	status := ColorGreen + "[OPERATIONAL]" + ColorReset
	
	fmt.Println()
	fmt.Println(ColorCyan + "╭─────────────── GHOST STATS ───────────────╮" + ColorReset)
	fmt.Printf(ColorCyan+"│"+ColorReset+" Files in shadow: "+ColorYellow+Bold+"%-18d"+ColorReset+ColorCyan+"│"+ColorReset+"\n", fileCount)
	fmt.Printf(ColorCyan+"│"+ColorReset+" Total size: "+ColorPurple+"%-25s"+ColorReset+ColorCyan+"│"+ColorReset+"\n", size)
	fmt.Printf(ColorCyan+"│"+ColorReset+" System status: %-26s"+ColorCyan+"│"+ColorReset+"\n", status)
	fmt.Println(ColorCyan + "╰────────────────────────────────────────────╯" + ColorReset)
	fmt.Println()
}

// GetPrompt retourne le prompt stylisé
func GetPrompt(fileCount int, totalSize int64) string {
	size := formatSize(totalSize)
	return fmt.Sprintf(ColorGreen+Bold+"[GHOST@system]"+ColorReset+ColorGray+"~"+ColorReset+" "+
		ColorGray+"[●] Files: "+ColorYellow+"%d"+ColorReset+ColorGray+" | Size: "+ColorPurple+"%s"+ColorReset+"\n"+
		ColorCyan+Bold+"└─>"+ColorReset+" ", fileCount, size)
}

// ShowCleanWarning affiche l'avertissement de nettoyage
func ShowCleanWarning() bool {
	fmt.Println()
	fmt.Println(ColorRed + Bold + "╔═══════════════════════════════════════════════════════════╗" + ColorReset)
	fmt.Println(ColorRed + Bold + "║        [!] WARNING: PERMANENT DELETION PROTOCOL           ║" + ColorReset)
	fmt.Println(ColorRed + Bold + "╚═══════════════════════════════════════════════════════════╝" + ColorReset)
	fmt.Println()
	fmt.Println(ColorYellow + "    All ghost files will be " + Bold + "ERASED" + ColorReset + ColorYellow + " from existence." + ColorReset)
	fmt.Println(ColorYellow + "    This action " + Bold + "CANNOT" + ColorReset + ColorYellow + " be undone." + ColorReset)
	fmt.Println()
	fmt.Print(ColorWhite + "    Type " + ColorGreen + Bold + "'CONFIRM'" + ColorReset + ColorWhite + " to proceed or any key to abort: " + ColorReset)
	
	var response string
	fmt.Scanln(&response)
	
	return strings.ToUpper(strings.TrimSpace(response)) == "CONFIRM"
}

// ShowGlitchEffect affiche un effet de glitch (easter egg)
func ShowGlitchEffect(text string) {
	glitched := ""
	for _, char := range text {
		if rand.Float32() < 0.3 {
			glitched += string(char) + "\u0336" // Ajoute une barre
		} else {
			glitched += string(char)
		}
	}
	fmt.Println(ColorGray + glitched + ColorReset)
}

// ClearScreen nettoie l'écran (compatible Windows et Unix)
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// ShowExitMessage affiche le message de sortie
func ShowExitMessage() {
	fmt.Println()
	fmt.Println(ColorGray + Dim + "    Disconnecting from ghost protocol..." + ColorReset)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(ColorCyan + "    [" + ColorGreen + "✓" + ColorCyan + "] Session terminated. Until next time, operator." + ColorReset)
	fmt.Println()
}