package src

import (
	"fmt"
	"strings"
	"unicode"
)

func afficherTitre() {
	fmt.Print(dore)
	fmt.Println(`
████████╗ ██████╗ ██╗    ██╗███████╗██████╗
╚══██╔══╝██╔═══██╗██║    ██║██╔════╝██╔══██╗
   ██║   ██║   ██║██║ █╗ ██║█████╗  ██████╔╝
   ██║   ██║   ██║██║███╗██║██╔══╝  ██╔══██╗
   ██║   ╚██████╔╝╚███╔███╔╝███████╗██║  ██║
   ╚═╝    ╚═════╝  ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝`)

	fmt.Print(blanc)
	fmt.Println("\n                         OF")

	fmt.Print(crimson)
	fmt.Println(`
██████╗ ██╗   ██╗ █████╗ ██╗     ██╗████████╗██╗   ██╗
██╔══██╗██║   ██║██╔══██╗██║     ██║╚══██╔══╝╚██╗ ██╔╝
██║  ██║██║   ██║███████║██║     ██║   ██║    ╚████╔╝
██║  ██║██║   ██║██╔══██║██║     ██║   ██║     ╚██╔╝
██████╔╝╚██████╔╝██║  ██║███████╗██║   ██║      ██║
╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝   ╚═╝      ╚═╝`)

	fmt.Print(reset)
	fmt.Println()
}

func (c *Character) MenuInitCharacter() {
	afficherTitre()
	titreEcran("CRÉATION DU PERSONNAGE")

	for {
		fmt.Print("Nom de votre aventurier : ")

		var nom string
		if _, err := fmt.Scan(&nom); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		valide := true
		for _, lettre := range nom {
			if !unicode.IsLetter(lettre) {
				valide = false
				break
			}
		}

		if !valide {
			messageErreur("Utilisez uniquement des lettres.")
			continue
		}

		lettres := []rune(strings.ToLower(nom))
		lettres[0] = unicode.ToUpper(lettres[0])
		c.Name = string(lettres)
		break
	}

	for {
		titreEcran("CHOISISSEZ VOTRE CLASSE")
		option("1", "Sparte • 60 PV max • 10 dégâts")
		option("2", "Athénien • 50 PV max • 15 dégâts")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch choix {
		case 1:
			c.Classe = "Sparte"
		case 2:
			c.Classe = "Athénien"
		default:
			messageErreur("Choisissez 1 ou 2.")
			continue
		}
		break
	}

	c.initCharacter()

	titreEcran("VOTRE AVENTURE COMMENCE")
	ligneInfo("Nom", c.Name)
	ligneInfo("Classe", c.Classe)
	ligneInfo("PV", fmt.Sprintf("%d/%d", c.Pv, c.PvMax))

	c.Menutest()
}

func (c *Character) MenuPrincipal() {
	for {
		titreEcran("PERSONNAGE")
		ligneInfo(c.Name+" • "+c.Level, fmt.Sprintf("%d/%d PV", c.Pv, c.PvMax))
		option("1", "Voir la fiche")
		option("2", "Ouvrir l'inventaire")
		option("3", "Aller au marché")
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
			c.displayInfo()
		case 2:
			c.accessInventory()
		case 3:
			c.MarketMenu()
		default:
			messageErreur("Choix invalide.")
		}
	}
}

func (c *Character) Menutest() {
	for {
		titreEcran("CARTE D'ATHÈNES")
		option("1", "Le Marché")
		option("2", "La Tour")
		option("3", "Les Maisons des Dieux")
		option("4", "Menu du personnage")
		option("0", "Quitter le jeu")
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
			c.MarketMenu()
		case 2:
			c.MenuTour()
		case 3:
			c.MenuMaisonsDieux()
		case 4:
			c.MenuPrincipal()
		default:
			messageErreur("Choix invalide.")
		}
	}
}