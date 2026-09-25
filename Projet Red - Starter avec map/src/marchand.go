package src

import (
	"fmt"
	"strings"
)

func (c *Character) MarketMenu() {
	for {
		titreEcran("MARCHÉ D'ATHÈNES")
		option("1", "Parler au marchand")
		option("2", "Parler au forgeron")
		option("0", "Retour à la carte")
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
			c.MenuMarchand()
		case 2:
			c.MenuForgeron()
		default:
			messageErreur("Choix invalide.")
		}
	}
}

func afficherArticle(numero, nom, prix string) {
	fmt.Printf(
		"  %s[%s]%s %s%-26s%s %s%s%s\n",
		dore, numero, reset,
		bleu, nom, reset,
		blanc, prix, reset,
	)
}

func (c *Character) afficherBourse() {
	mot := "pièces"
	if c.Money == 1 {
		mot = "pièce"
	}

	ligneInfo("Bourse", fmt.Sprintf("%d %s d'or", c.Money, mot))
}

func (c *Character) DisplayMarket() {
	titreEcran("ÉTAL DU MARCHAND")
	c.afficherBourse()

	sousTitre("POTIONS")
	afficherArticle("1", "Ambroisie", "1 pièce")
	afficherArticle("2", "Poison", "2 pièces")
	afficherArticle("3", "Potion de Pâques", "10 pièces • une fois")

	sousTitre("ARMES")
	afficherArticle("4", "Dagues d'Assassin", "3 pièces")
	afficherArticle("5", "Glaive", "3 pièces")
	afficherArticle("6", "Marteau de Guerre", "3 pièces")
	afficherArticle("7", "Hache de Viking", "3 pièces")

	sousTitre("MATÉRIAUX ET SERVICES")
	afficherArticle("8", "Fer", "1 pièce")
	afficherArticle("9", "Cuir", "1 pièce")
	afficherArticle("10", "Inventaire +10 places", "30 pièces • 3 fois max")

	if !c.PotionGratuiteRecuperee {
		messageSucces("Première ambroisie offerte.")
	}

	option("M", "Revoir l'étal")
	option("0", "Retourner au marché")
}

func (c *Character) MenuMarchand() {
	c.DisplayMarket()

	for {
		fmt.Print("\nArticle (M : étal, 0 : retour) : ")

		var choix string
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch strings.ToUpper(choix) {
		case "0":
			return
		case "M":
			c.DisplayMarket()
			continue

		case "1":
			if !c.PotionGratuiteRecuperee {
				if !c.CheckPlace() {
					messageErreur("Inventaire plein.")
					continue
				}
				c.addinventory(PotionSoin, 1)
				c.PotionGratuiteRecuperee = true
				messageSucces("Le marchand vous offre une potion d'ambroisie.")
			} else {
				c.acheterObjet(PotionSoin, 1)
			}

		case "2":
			c.acheterObjet(PotionPoison, 2)

		case "3":
			if c.PotionPaqueAchetee {
				messageErreur("La potion de Pâques ne peut être achetée qu'une fois.")
				continue
			}
			if c.acheterObjet(PotionPaque, 10) {
				c.PotionPaqueAchetee = true
			}

		case "4":
			c.acheterObjet("Les Dagues de l'Assassin", 3)
		case "5":
			c.acheterObjet("Le Glaive du Légionnaire", 3)
		case "6":
			c.acheterObjet("Le Marteau de Guerre du Martelier", 3)
		case "7":
			c.acheterObjet("La Hache de Viking", 3)
		case "8":
			c.acheterObjet("Fer", 1)
		case "9":
			c.acheterObjet("Cuir", 1)
		case "10":
			c.acheterAmeliorationInventaire()

		default:
			messageErreur("Choix invalide.")
			continue
		}

		c.afficherBourse()
	}
}

func (c *Character) acheterObjet(nom string, prix int) bool {
	if !c.CheckPlace() {
		messageErreur("Inventaire plein.")
		return false
	}

	if !c.Moni(prix) {
		return false
	}

	c.addinventory(nom, 1)
	messageSucces("Vous obtenez : " + nom)
	return true
}

func (c *Character) acheterAmeliorationInventaire() {
	if c.UpgradeCount >= 3 {
		messageErreur("Les trois améliorations ont déjà été achetées.")
		return
	}

	if !c.Moni(30) {
		return
	}

	c.UpgradeInventorySlot()
	messageSucces(fmt.Sprintf(
		"Inventaire amélioré : %d places.",
		c.LimitInventaire,
	))
}