package src

import "fmt"

func (c *Character) MerchantMenu() {
	// afficher les objets du marchand
	// demander le choix
	// ajouter l'objet à l'inventaire
	var chose int
	for true {
		fmt.Println("=== Marché ===")
		fmt.Println("\t 1 - Marchand")
		fmt.Println("\t 2 - Forgeron")
		fmt.Println("\t 0 - Retour à la carte")

		fmt.Print("Votre choix : ")
		fmt.Scan(&chose)

		switch chose {
		case 1:
			return
		case 2:
			fmt.Println("Vous vous dirigez vers la Tour.")
		case 0:
			fmt.Println("alaide")
			break
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
