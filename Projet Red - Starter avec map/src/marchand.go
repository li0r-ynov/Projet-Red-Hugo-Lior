package src

import "fmt"

func (c *Character) DisplayMarket() {
	fmt.Println("=== Marché ===")
	fmt.Println("\t 1 - La potion d'ambroisie (3 pièces d'or)")
	fmt.Println("\t 2 - La potion de poison (5 pièces d'or)")
	fmt.Println("\t 3 - Les Dagues de l'Assassin' (5 pièces d'or)")
	fmt.Println("\t 4 - Le Glaives du Légionnaire(5 pièces d'or)")
	fmt.Println("\t 5 - Le Marteau de guerre du Martelier (5 pièces d'or)")
	fmt.Println("\t 6 - Le Hache de Vikings (5 pièces d'or)")

	fmt.Println("\t potion de pâques")

	fmt.Println("\t 0 - Retour à la carte")
	fmt.Printf("\t Pièces d'or : %d\n", c.Money)
	fmt.Println("------------------------------")
	fmt.Println("Votre choix ?")
}

/* func (c *Character) count(prix int) {
	if c.Moni < prix {
		fmt.Println("vous n'avez pas assez d epièces d'or.")
		return
	}
	fmt.Printf("il vous reste %d pièces d'or.%n", c.Moni)
	c.Moni(10)
}
*/
func (c *Character) MarketMenu() {
	for true {
		c.DisplayMarket()
		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 1:
			if !c.PotionGratuiteRecuperee {
				fmt.Println("\t Le marchand vous offre une Potion d'ambroisie")
				c.addinventory("Potion d'ambroisie", 1)
				c.PotionGratuiteRecuperee = true
			} else {
				fmt.Println("\t Vous venez d'acheter une potion d'ambroisie")
				c.addinventory("Potion d'ambroisie", 1)
				c.Moni(3)
			}
		case 2:
			fmt.Println("\t Vous venez d'acheter une potion de poison hahaha !")
			c.addinventory("Potion de poison", 1)
			c.Moni(5)
		case 3:
			fmt.Println("\t Vous venez d'acheter les Dagues de l'Assassin !")
			c.addinventory("Les Dagues de l'Assassin", 1)
			c.Moni(5)
		case 4:
			fmt.Println("\t Vous venez d'acheter la Hache de Viking !")
			c.addinventory("La Hache de Viking", 1)
			c.Moni(5)
		case 5:
			fmt.Println("\t Vous venez d'acheter le Marteau de Guerre du Martellier!")
			c.addinventory("Le Marteau de Guerre du Martellier", 1)
			c.Moni(5)
		case 6:
			fmt.Println("\t Vous venez d'acheter le Glaive du Légionnaire ")
			c.addinventory("Le Glaive du Légionnaire", 1)
			c.Moni(5)
		case 0:
			fmt.Println("alaide")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}

}

/* func cost int (){

} */

/* fmt.Println("\t Vous venez d'acheter une potion de pâques ") */

/* func marketplace() {



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
