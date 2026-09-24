package src

import "fmt"

func (c *Character) TowerTravelDisplay(
	text, textOption1, textOption2, textOption3, textOption4, textOption5 string,
) int {
	fmt.Println(text)
	fmt.Println("\t 1 - ", textOption1)
	fmt.Println("\t 2 - ", textOption2)
	fmt.Println("\t 3 - ", textOption3)
	fmt.Println("\t 4 - ", textOption4)
	fmt.Println("\t 0 - ", textOption5)
	fmt.Print("Votre choix : ")

	var loreChoice int
	fmt.Scan(&loreChoice)
	return loreChoice
}

func (c *Character) MenuTour() {
	for c.DirectionTour == "" {
		fmt.Println("\n=== Entrée de la tour ===")
		fmt.Println("Choisissez votre voie. Ce choix est définitif.")
		fmt.Println("1 - Monter")
		fmt.Println("2 - Descendre")
		fmt.Println("0 - Retourner en ville")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			c.DirectionTour = "Monter"
		case 2:
			c.DirectionTour = "Descendre"
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}

	for {
		if c.EtageTour > 5 {
			fmt.Println("Vous avez terminé les cinq combats de votre voie.")
			return
		}

		nomAdversaire := adversaireEtage(c.DirectionTour, c.EtageTour)
		adversaire, existe := oppsdef[nomAdversaire]
		if !existe {
			fmt.Printf("Adversaire introuvable : %s\n", nomAdversaire)
			return
		}

		fmt.Printf("\n=== %s : étage %d ===\n", c.DirectionTour, c.EtageTour)
		fmt.Printf("Adversaire : %s\n", adversaire.Name)
		fmt.Println("1 - Commencer le combat")
		fmt.Println("0 - Retourner en ville")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			return

		case 1:
			if c.Pv <= 0 {
				fmt.Println("Vous devez récupérer des PV avant de combattre.")
				return
			}

			etageVaincu := c.EtageTour

			if !c.combat(adversaire) {
				fmt.Println("Vous pourrez retenter cet étage plus tard.")
				return
			}

			c.recompenseEtage(etageVaincu)
			c.EtageTour++

			fmt.Println("Étage terminé ! Vous pouvez continuer ou retourner en ville.")

		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func adversaireEtage(direction string, etage int) string {
	switch direction {
	case "Monter":
		switch etage {
		case 1:
			return "Combattant"
		case 2:
			return "Guerrier"
		case 3:
			return "Héro"
		case 4:
			return "Demi-Dieu"
		case 5:
			return "Zeus"
		}

	case "Descendre":
		switch etage {
		case 1:
			return "Bandit"
		case 2:
			return "Mercenaire"
		case 3:
			return "Démon"
		case 4:
			return "Demi-Dieu"
		case 5:
			return "Hadès"
		}
	}

	return ""
}

func (c *Character) recompenseEtage(etage int) {
	seuils := []int{0, 100, 300, 700, 1500, 3100}
	renommeeVoulue := seuils[etage]

	if c.DirectionTour == "Descendre" {
		renommeeVoulue = -renommeeVoulue
	}

	// ChangeRenown ajoute une quantité : on calcule la différence
	// pour atteindre exactement le seuil voulu.
	gain := renommeeVoulue - c.Renown
	c.ChangeRenown(gain)

	fmt.Printf("Votre titre est maintenant : %s.\n", c.Level)
	fmt.Printf("Votre renommée est maintenant de %d.\n", c.Renown)

	switch etage {
	case 3:
		c.donnerButin("Plume divine", 3)

	case 4:
		c.donnerButin("Essence divine", 2)
	}
}

func (c *Character) donnerButin(nom string, quantite int) {
	totalObjets := 0
	for _, nombre := range c.Inventaire {
		totalObjets += nombre
	}

	placeDisponible := c.LimitInventaire - totalObjets
	if placeDisponible < quantite {
		quantite = placeDisponible
	}

	if quantite <= 0 {
		fmt.Printf("Votre inventaire est plein : impossible de récupérer %s.\n", nom)
		return
	}

	c.addinventory(nom, quantite)
	fmt.Printf("Butin obtenu : %s x%d.\n", nom, quantite)
}