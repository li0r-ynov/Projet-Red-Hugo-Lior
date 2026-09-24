package src

import "fmt"

// MarketMenu est l'entrée du marché.
func (c *Character) MarketMenu() {
	for {
		fmt.Println("\n=== Marché ===")
		fmt.Println("1 - Parler au marchand")
		fmt.Println("2 - Parler au forgeron")
		fmt.Println("0 - Retourner à la carte")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("Saisie invalide.")
			return
		}

		switch choix {
		case 1:
			c.MenuMarchand()
		case 2:
			c.MenuForgeron()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) DisplayMarket() {
	fmt.Println("\n=== Marchand ===")
	fmt.Println("1 - Potion d'ambroisie (1 pièce d'or, première offerte)")
	fmt.Println("2 - Potion de poison (2 pièces d'or)")
	fmt.Println("3 - Les Dagues de l'Assassin (3 pièces d'or)")
	fmt.Println("4 - Le Glaive du Légionnaire (3 pièces d'or)")
	fmt.Println("5 - Le Marteau de Guerre du Martelier (3 pièces d'or)")
	fmt.Println("6 - La Hache de Viking (3 pièces d'or)")
	fmt.Println("7 - Fer (1 pièce d'or)")
	fmt.Println("8 - Cuir (1 pièce d'or)")
	fmt.Println("0 - Retourner au marché")
	fmt.Printf("Pièces d'or : %d\n", c.Money)
	fmt.Print("Votre choix : ")
}

func (c *Character) MenuMarchand() {
	for {
		c.DisplayMarket()

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("Saisie invalide.")
			return
		}

		switch choix {
		case 0:
			return

		case 1:
			if !c.PotionGratuiteRecuperee {
				if !c.CheckPlace() {
					fmt.Println("Votre inventaire est plein.")
					continue
				}

				c.addinventory(PotionSoin, 1)
				c.PotionGratuiteRecuperee = true
				fmt.Println("Le marchand vous offre une potion d'ambroisie.")
			} else {
				c.acheterObjet(PotionSoin, 1)
			}

		case 2:
			c.acheterObjet(PotionPoison, 2)
		case 3:
			c.acheterObjet("Les Dagues de l'Assassin", 3)
		case 4:
			c.acheterObjet("Le Glaive du Légionnaire", 3)
		case 5:
			c.acheterObjet("Le Marteau de Guerre du Martelier", 3)
		case 6:
			c.acheterObjet("La Hache de Viking", 3)
		case 7:
			c.acheterObjet("Fer", 1)
		case 8:
			c.acheterObjet("Cuir", 1)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (c *Character) acheterObjet(nom string, prix int) {
	if !c.CheckPlace() {
		fmt.Println("Votre inventaire est plein.")
		return
	}

	if c.Money < prix {
		fmt.Println("Vous n'avez pas assez de pièces d'or.")
		return
	}

	c.Money -= prix
	c.addinventory(nom, 1)
	fmt.Printf("Vous obtenez : %s.\n", nom)
}