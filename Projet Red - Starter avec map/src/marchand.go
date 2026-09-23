package src

import "fmt"

func DisplayMarket() {
	fmt.Println("=== Marché ===")
	fmt.Println("\t 1 - potion d'ambroisie (10 pièces d'or)")
	fmt.Println("\t 2 - potion de poison (5 pièces d'or)")
	fmt.Println("\t potion de pâques")
	fmt.Println("\t Le marchand vous offre une Potion d'ambroisie")
	fmt.Println("\t Vous venez d'acheter une potion d'ambroisie")
	fmt.Println("\t vous venez d'acheter une potion de poison hahaha !")
	fmt.Println("\t Vous venez d'acheter une potion de pâques ")
	fmt.Println("\t 0 - Retour à la carte")
	fmt.Printf("\t Pièces d'or : %d\n", c.money)
	fmt.Println("------------------------------")
	fmt.Println("Votre choix ?")
}

func (c *Character) MarketMenu() {
	for true {
		DisplayMarket()
		var chose int
		fmt.Scan(&chose)

		if chose != 0 && !c.CheckPlace() {
			fmt.Println("\nImpossible pas de place dans l'inventaire....")
			continue
		}

		switch chose {
		case 1:
			if c.Money(10) {
				c.Inventaire["Potion d'ambroisie"] += 1
				fmt.Println("-10 pièces d'or, vous avez acheté une potion d'ambroisie")
			}

		case 2:
			fmt.Println("Vous vous dirigez vers la Tour.")
		case 0:
			fmt.Println("alaide")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}

}

/* func marketplace() {

	if !c.PotionGratuiteRecuperee {
		fmt.Println("\t1 - Potion d'ambroisie - GRATUITE")
	} else {
		fmt.Println("\t1 - Potion d'ambroisie - 3 pièces d'or")
	}

	fmt.Println("\t0 - Retour")

	fmt.Print("Votre choix : ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if !c.PotionGratuiteRecuperee {
			c.Inventaire[PotionSoin]++
			c.PotionGratuiteRecuperee = true

			fmt.Println("Le marchand vous offre une Potion d'ambroisie !")
		} else {
			fmt.Println("La Potion d'ambroisie coûte maintenant 3 pièces d'or.")
			// Le système d'argent sera ajouté avec la tâche 13/14.
		}

	case 0:
		return

	default:
		fmt.Println("Choix invalide.")
	}
} */
