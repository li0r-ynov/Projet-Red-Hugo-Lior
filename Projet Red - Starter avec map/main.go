package main

import (
	"fmt"
)

// main lance le jeu : création du personnage puis boucle sur le menu principal.
func main() {
	var player Character
	player.initCharacter("athénien", "civil")
	
	for true {
		fmt.Println("=== Menu Principal ===")
		fmt.Println("\t 1 - Afficher les informations du personnage")
		fmt.Println("\t 2 - Accéder à l'inventaire")
		fmt.Println("\t 0 - Quitter le jeu")

		fmt.Print("Votre choix : ")
		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			return
		case 1:
			player.displaylnfo()
		case 2:
			player.accessInventory()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
