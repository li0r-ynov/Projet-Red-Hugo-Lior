package main

import "fmt"

func (c *Character) TowerTravelDisplay(text, textOption1, textOption2 string) int {
	fmt.Println(text)
	fmt.Println("\t 1 - ", textOption1)
	fmt.Println("\t 2 - ", textOption2)

	fmt.Print("Votre choix : ")
	var LoreChoice int
	fmt.Scan(&LoreChoice)
}

/*
func (c *Character) accessInventory() {
	for true {
		fmt.Println("=== Inventaire ===")
		for itemName, itemQuantity := range c.Inventaire {
			fmt.Printf("\t - %s x%d\n", itemName, itemQuantity)
		}
		fmt.Println("\t 1 - Utiliser une potion de soin")
		fmt.Println("\t 0 - Retour au menu principal")

		fmt.Print("Votre choix : ")
		var test int
		fmt.Scan(&test)

		switch test {
		case 0:
			return
		case 1:
			c.takePot()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}

}

func (c *Character) takePot() {
	potQuantity, potCheck := c.Inventaire[PotionPv]
	if !potCheck {
		fmt.Println("Vous n'avez aucune potion dans votre inventaire.")
		return
	}

	if potQuantity <= 0 {
		fmt.Println("Il ne vous reste plus de potion disponible.")
		return
	}

	c.Pv += 50
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}
	c.Inventaire[PotionPv]--
	if c.Inventaire[PotionPv] <= 0 {
		delete(c.Inventaire, PotionPv)
	}

	fmt.Printf("Potion utilisée (-1) ! Vous avez maintenant %d/%d PV.\n", c.Pv, c.PvMax)
}

func main() {
	var player Character
	player.initCharacter("Cyril", "mentor")
	for true {
		fmt.Println("=== Menu Principal ===")
		fmt.Println("\t 1 - Afficher les informations du personnage")
		fmt.Println("\t 2 - Accéder à l'inventaire")
		fmt.Println("\t 0 - Quitter le jeu")

		fmt.Print("Votre choix : ")
		var chose int
		fmt.Scan(&chose)
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
