package main

import (
	"fmt"
)

// main lance le jeu : création du personnage puis boucle sur le menu principal.
func (c *Character)MenuPrincipale() {
	var player Character
	player.initCharacter("athénien", "civil")
	
	for true {
		fmt.Println("=== Menu Principal ===")
		fmt.Println("\t 1 - Afficher les informations du personnage")
		fmt.Println("\t 2 - Accéder à l'inventaire")
		fmt.Println("\t 0 - Retour à la carte")

		fmt.Print("Votre choix : ")
		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			return
		case 1:
			c.displaylnfo()
		case 2:
			c.accessInventory()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

// main lance le jeu : création du personnage puis boucle sur le menu principal.
func main() {
	var player Character

	player.initCharacter("athénien", "civil")
	for true {
		step1 := player.TowerTravelDisplay(
			"Où souhaitez-vous vous rendre ?",
			"Le Marché",
			"La Tour",
			"Les Maisons des Dieux",
			"Menu Principal",
			"quitter le jeu",
		)

		switch step1 {
		case 1:
			fmt.Println("Vous vous dirigez vers le Marché.")

		case 2:
			fmt.Println("Vous vous dirigez vers la Tour.")

		case 3:
			fmt.Println("Vous vous dirigez vers les Maisons des Dieux.")

		case 4:
			player.MenuPrincipale()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}

}
