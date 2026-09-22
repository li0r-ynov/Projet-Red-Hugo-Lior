package main

import (
	"fmt"
)

func (c *Character) accessInventory() {
	for {
		fmt.Println("=== Inventaire ===")
		for itemName, itemQuantity := range c.Inventaire {
			fmt.Printf("\t - %s x%d\n", itemName, itemQuantity)
		}
		fmt.Println("\t 1 - Utiliser une potion d'ambroisie")
		fmt.Println("\t 2 - Utiliser une potion de poison")
		fmt.Println("\t 3 - Utiliser une potion de pâques")
		fmt.Println("\t 0 - Retour au menu principal")

		fmt.Print("Votre choix : ")
		var test int
		fmt.Scan(&test)

		switch test {
		case 0:
			return
		case 1, 2, 3:
			c.takePot(test)
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) Moni(cost int){
	var money int 
	if money <  cost{
		fmt.Println("T'as pas les tales clochard")
	} else {
		money -= cost
		fmt.Println("Merci pour vos achats")
	}
}
