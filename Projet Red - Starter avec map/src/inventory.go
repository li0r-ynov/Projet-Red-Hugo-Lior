package src

import (
	"fmt"
	"sort"
)

func (c *Character) accessInventory() {
	for {
		titreEcran("INVENTAIRE")

		total := 0
		for _, quantite := range c.Inventaire {
			total += quantite
		}

		fmt.Printf("  %sPlaces%s : %s%d/%d%s\n",
			bleu, reset, blanc, total, c.LimitInventaire, reset)

		c.afficherCategorieInventaire("POTIONS", "potion")
		c.afficherCategorieInventaire("ARMES", "arme")
		c.afficherCategorieInventaire("ARMURES", "armure")
		c.afficherCategorieInventaire("MATÉRIAUX", "materiau")

		fmt.Println()
		separateur()
		option("1", "Utiliser une potion d'ambroisie")
		option("2", "Utiliser une potion de poison")
		option("3", "Utiliser une potion de pâques")
		option("4", "Équiper une arme")
		option("5", "Équiper une armure")
		option("0", "Retour")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch choix {
		case 0:
			return
		case 1, 2, 3:
			c.takePot(choix)
		case 4:
			c.weaponEquipmentMenu()
		case 5:
			c.armorEquipmentMenu()
		default:
			messageErreur("Choix invalide.")
		}
	}
}

func (c *Character) afficherCategorieInventaire(titre, categorie string) {
	var objets []string

	for nom, quantite := range c.Inventaire {
		if quantite > 0 && categorieObjet(nom) == categorie {
			objets = append(objets, nom)
		}
	}

	if len(objets) == 0 {
		return
	}

	sort.Strings(objets)
	sousTitre(titre)

	for _, nom := range objets {
		fmt.Printf("  %s• %-30s%s %sx%d%s\n",
			bleu, nom, reset,
			blanc, c.Inventaire[nom], reset)
	}
}

func categorieObjet(nom string) string {
	switch nom {
	case PotionSoin, PotionPoison, PotionPaques:
		return "potion"
	case "Fer", "Cuir", "Plume divine", "Essence divine":
		return "materiau"
	}

	if _, existe := offgear[nom]; existe {
		return "arme"
	}

	if _, existe := defgear[nom]; existe {
		return "armure"
	}

	return "materiau"
}

func (c *Character) armorEquipmentMenu() {
	var armures []string

	for nom, quantite := range c.Inventaire {
		_, existe := defgear[nom]
		if existe && quantite > 0 {
			armures = append(armures, nom)
		}
	}

	sort.Strings(armures)
	titreEcran("ÉQUIPER UNE ARMURE")

	if len(armures) == 0 {
		fmt.Println("  Vous ne possédez aucune armure.")
		return
	}

	for i, nom := range armures {
		fmt.Printf("  %s[%d]%s %s%s%s\n",
			dore, i+1, reset, bleu, nom, reset)
	}

	option("0", "Retour")
	fmt.Print("Votre choix : ")

	var choix int
	if _, err := fmt.Scan(&choix); err != nil {
		messageErreur("Saisie invalide.")
		return
	}

	if choix == 0 {
		return
	}

	if choix < 1 || choix > len(armures) {
		messageErreur("Choix invalide.")
		return
	}

	c.equipArmor(armures[choix-1])
}

func (c *Character) Moni(cost int) bool {
	if c.Money < cost {
		messageErreur("Vous n'avez pas assez de pièces d'or.")
		return false
	}

	c.Money -= cost
	messageSucces("Merci pour votre achat.")
	return true
}

func (c *Character) CheckPlace() bool {
	total := 0

	for _, quantite := range c.Inventaire {
		total += quantite
	}

	return total < c.LimitInventaire
}

func (c *Character) addinventory(item string, quantity int) {
	if c.Inventaire == nil {
		c.Inventaire = make(map[string]int)
	}

	c.Inventaire[item] += quantity
}