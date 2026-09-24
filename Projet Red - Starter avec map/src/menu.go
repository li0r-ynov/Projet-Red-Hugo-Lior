package src

import (
	"fmt"
	"strings"
	"unicode"
)

func (c *Character) MenuInitCharacter() {
	var nom string
	fmt.Print("Quel est votre nom jeune aventurier :\n")
	fmt.Scan(&nom)

	//met le nom au bon format
	for _, caractere := range nom {
		if !unicode.IsLetter(caractere) {
			fmt.Println("nom invalide :  utilisez uniquement des lettres.")
			c.MenuInitCharacter()
			return
		}
	}
	nom = strings.ToLower(nom)
	premierelettre := nom[:1]
	reste := nom[1:]
	premierelettre = strings.ToUpper(premierelettre)
	nom = premierelettre + reste
	c.Name = nom
	fmt.Println("nom final :", c.Name)

	var choiceType int
	fmt.Println("1 - Sparte")
	fmt.Println("2 - Athénien")
	fmt.Scan(&choiceType)

	switch choiceType {

	case 1:
		c.Classe = "Sparte"
	case 2:
		c.Classe = "Athénien"
	}
	c.initCharacter()
	// fmt.Println(choiceType)
	c.MenuPrincipal()
}

func (c *Character) MenuPrincipal() {

	for true {
		fmt.Println("=== Menu Principal ===")
		fmt.Println("\t 1 - Afficher les informations du personnage")
		fmt.Println("\t 2 - Accéder à l'inventaire")
		fmt.Println("\t 3 - Accéder au shop")
		fmt.Println("\t 0 - sortir du jeu")

		fmt.Print("Votre choix : ")
		var chose int
		fmt.Scan(&chose)

		switch chose {
		case 0:
			return
		case 1:
			c.displayInfo()
		case 2:
			c.accessInventory()
		case 3:
			c.MarketMenu()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}

func (c *Character) Menutest() {
	c.initCharacter()
	for true {
		step1 := c.TowerTravelDisplay(
			"Où souhaitez-vous vous rendre ?",
			"Le Marché",
			"La Tour",
			"Les Maisons des Dieux",
			"Menu Principal",
			"quitter le jeu",
		)

		switch step1 {
		case 1:
			fmt.Println("Vous vous dirigez vers le Marché.")
			c.MarketMenu()
		case 2:
			fmt.Println("Vous vous dirigez vers la Tour.")

		case 3:
			fmt.Println("Vous vous dirigez vers les Maisons des Dieux.")

		case 4:
			c.MenuPrincipal()
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
