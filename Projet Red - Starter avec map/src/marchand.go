package src

import "fmt"

func (c *Character) DisplayMarket() {
	fmt.Println("=== Marché ===")
	fmt.Println("\t 1 - Potion d'Ambroisie (3 pièces d'or)")
	fmt.Println("\t 2 - Potion de Poison (5 pièces d'or)")
	fmt.Println("\t 3 - Potion de Pâque (5 pièces d'or)")
	fmt.Println("\t 4 - Dagues de l'Assassin (5 pièces d'or)")
	fmt.Println("\t 5 - Glaive du Légionnaire(5 pièces d'or)")
	fmt.Println("\t 6 - Marteau de guerre du Martelier (5 pièces d'or)")
	fmt.Println("\t 7 - Hache de Vikings (5 pièces d'or)")
	fmt.Println("\t 0 - Retour menu principal")
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

func acheter(c *Character, nomObjet string, prix int) {
	if c.addInventory(nomObjet, 1){
	fmt.Println("\t Vous venez d'acheter", nomObjet)
	c.Moni(prix)}
}

func (c *Character) MarketMenu() {
	for true {
		c.DisplayMarket()
		var chose int
		fmt.Scan(&chose)

		switch chose {

		case 1:
			if !c.PotionGratuiteRecuperee {
				fmt.Println("\t Le marchand vous offre une Potion d'ambroisie")
				c.addInventory("Potion d'ambroisie", 1)
				c.PotionGratuiteRecuperee = true
			} else {
				acheter(c, "Potion d'ambroisie", 3)
			}
		case 2:
			acheter(c, "Potion de poison", 5)

		case 3:
			acheter(c, "Potion de pâque", 5)

		case 4:
			acheter(c, "Les Dagues de l'Assasin", 5)

		case 5:
			acheter(c, "Le Glaive du Légionnaire", 5)

		case 6:
			acheter(c, "Le Marteau de Guerre du Martellier", 5)

		case 7:
			acheter(c, "La Hache de Vikings", 5)
		case 0:
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
