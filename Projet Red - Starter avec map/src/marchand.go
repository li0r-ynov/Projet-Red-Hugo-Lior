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
		option("0", "Retourner à la carte")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
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
			messageErreur("Choix invalide.")
		}
	}
}

func (c *Character) DisplayMarket() {
	titreEcran("ÉTAL DU MARCHAND")
	c.afficherBourse()

	sousTitre("POTIONS")
	afficherArticle("1", "🧪", "Ambroisie", "1 pièce")
	afficherArticle("2", "☠", "Poison", "2 pièces")

	sousTitre("ARMES")
	afficherArticle("3", "🗡", "Dagues d'Assassin", "3 pièces")
	afficherArticle("4", "⚔", "Glaive", "3 pièces")
	afficherArticle("5", "🔨", "Marteau de Guerre", "3 pièces")
	afficherArticle("6", "🪓", "Hache de Viking", "3 pièces")

	sousTitre("MATÉRIAUX")
	afficherArticle("7", "⛓", "Fer", "1 pièce")
	afficherArticle("8", "▰", "Cuir", "1 pièce")

	fmt.Println()
	separateur()

	if !c.PotionGratuiteRecuperee {
		fmt.Printf("  %s✦ Première ambroisie offerte !%s\n", crimson, reset)
	}

	option("M", "Revoir l'étal")
	option("0", "Retourner au marché")
}

func (c *Character) afficherBourse() {
	mot := "pièces"
	if c.Money == 1 {
		mot = "pièce"
	}

	fmt.Printf("  %sBourse%s : %s%d %s d'or%s\n",
		bleu, reset, blanc, c.Money, mot, reset)
}

func afficherArticle(numero, icone, nom, prix string) {
	fmt.Printf("  %s[%s]%s %s %s%-23s%s %s%s%s\n",
		dore, numero, reset,
		icone,
		bleu, nom, reset,
		blanc, prix, reset)
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

		case "1":
			if !c.PotionGratuiteRecuperee {
				if !c.CheckPlace() {
					messageErreur("Votre inventaire est plein.")
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
			c.acheterObjet("Les Dagues de l'Assassin", 3)
		case "4":
			c.acheterObjet("Le Glaive du Légionnaire", 3)
		case "5":
			c.acheterObjet("Le Marteau de Guerre du Martelier", 3)
		case "6":
			c.acheterObjet("La Hache de Viking", 3)
		case "7":
			c.acheterObjet("Fer", 1)
		case "8":
			c.acheterObjet("Cuir", 1)
		default:
			messageErreur("Choix invalide. Entrez un numéro, M ou 0.")
			continue
		}

		c.afficherBourse()
	}
}

func (c *Character) acheterObjet(nom string, prix int) {
	if !c.CheckPlace() {
		messageErreur("Votre inventaire est plein.")
		return
	}

	if c.Money < prix {
		messageErreur("Vous n'avez pas assez de pièces d'or.")
		return
	}

	c.Money -= prix
	c.addinventory(nom, 1)
	messageSucces("Vous obtenez : " + nom)
}