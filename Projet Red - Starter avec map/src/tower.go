package src

import "fmt"

func (c *Character) TowerTravelDisplay(
	text, textOption1, textOption2, textOption3, textOption4, textOption5 string,
) int {
	fmt.Println(text)
	option("1", textOption1)
	option("2", textOption2)
	option("3", textOption3)
	option("4", textOption4)
	option("0", textOption5)
	fmt.Print("Votre choix : ")

	var choix int
	fmt.Scan(&choix)
	return choix
}

func (c *Character) MenuTour() {
	for c.DirectionTour == "" {
		titreEcran("LA TOUR")
		fmt.Printf("  %sDeux voies s'ouvrent devant vous.%s\n", blanc, reset)
		fmt.Printf("  %sVotre choix sera définitif.%s\n", crimson, reset)
		separateur()
		option("1", "Monter vers l'Olympe")
		option("2", "Descendre vers les Enfers")
		option("0", "Retourner à Athènes")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch choix {
		case 1:
			c.DirectionTour = "Monter"
		case 2:
			c.DirectionTour = "Descendre"
		case 0:
			return
		default:
			messageErreur("Choix invalide.")
		}
	}

	for {
		if c.EtageTour > 5 {
			titreEcran("VOIE TERMINÉE")
			messageSucces("Vous avez remporté les cinq combats de votre voie.")
			return
		}

		nomAdversaire := adversaireEtage(c.DirectionTour, c.EtageTour)
		adversaire, existe := oppsdef[nomAdversaire]

		if !existe {
			messageErreur("Adversaire introuvable : " + nomAdversaire)
			return
		}

		titreEcran("PARCOURS DE LA TOUR")
		ligneTour("Voie", c.DirectionTour)
		ligneTour("Étage", fmt.Sprintf("%d/5", c.EtageTour))
		ligneTour("Adversaire", adversaire.Name)
		separateur()

		option("1", "Commencer le combat")
		option("0", "Retourner à Athènes")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch choix {
		case 0:
			return

		case 1:
			if c.Pv <= 0 {
				messageErreur("Vous devez récupérer des PV avant de combattre.")
				return
			}

			etageVaincu := c.EtageTour

			if !c.combat(adversaire) {
				messageErreur("Vous pourrez retenter cet étage plus tard.")
				return
			}

			c.recompenseEtage(etageVaincu)
			c.EtageTour++
			messageSucces("Étage terminé ! Continuez ou retournez en ville.")

		default:
			messageErreur("Choix invalide.")
		}
	}
}

func ligneTour(nom, valeur string) {
	fmt.Printf("  %s%-13s%s %s%s%s\n",
		bleu, nom, reset,
		blanc, valeur, reset)
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

	c.ChangeRenown(renommeeVoulue - c.Renown)

	sousTitre("RÉCOMPENSES")
	ligneTour("Titre", c.Level)
	ligneTour("Renommée", fmt.Sprintf("%d", c.Renown))

	switch etage {
	case 3:
		// Héro sur la voie montante, Démon sur la voie descendante.
		c.donnerButin("Plume divine", 2)

	case 4:
		// Demi-Dieu sur les deux voies.
		c.donnerButin("Essence divine", 1)
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
		messageErreur("Inventaire plein : impossible de récupérer " + nom + ".")
		return
	}

	c.addinventory(nom, quantite)
	messageSucces(fmt.Sprintf("Butin obtenu : %s x%d", nom, quantite))
}