package src

import (
	"fmt"
	"sort"
)

func (c *Character) accessInventory() {
	for {
		titreEcran("INVENTAIRE")
		ligneInfo("Places", fmt.Sprintf("%d/%d", c.nombreObjets(), c.LimitInventaire))

		for _, categorie := range []string{
			"POTIONS", "ARMES", "ARMURES", "MATÉRIAUX", "AUTRES",
		} {
			c.afficherCategorieInventaire(categorie)
		}

		fmt.Println()
		separateur()
		option("1", "Utiliser une potion d'ambroisie")
		option("2", "Utiliser une potion de poison")
		option("3", "Utiliser une potion de Pâques")
		option("4", "Équiper ou retirer une arme")
		option("5", "Équiper ou retirer une armure")
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

func (c *Character) Moni(cost int) bool {
	if c.Money < cost {
		messageErreur("T'as pas les tales CLOCHARD !!!")
		return false
	}

	c.Money -= cost
	return true
}

func (c *Character) CheckPlace() bool {
	return c.nombreObjets() < c.LimitInventaire
}

func (c *Character) addinventory(item string, quantity int) {
	if c.Inventaire == nil {
		c.Inventaire = make(map[string]int)
	}

	c.Inventaire[item] += quantity
}

func (c *Character) removeInventory(item string, quantity int) bool {
	if quantity <= 0 || c.Inventaire[item] < quantity {
		return false
	}

	c.Inventaire[item] -= quantity
	if c.Inventaire[item] == 0 {
		delete(c.Inventaire, item)
	}

	return true
}

func (c *Character) nombreObjets() int {
	total := 0
	for _, quantity := range c.Inventaire {
		total += quantity
	}
	return total
}

func (c *Character) UpgradeInventorySlot() bool {
	if c.UpgradeCount >= 3 {
		return false
	}

	c.LimitInventaire += 10
	c.UpgradeCount++
	return true
}

func categorieObjet(nom string) string {
	switch nom {
	case PotionSoin, PotionPoison, PotionPaque:
		return "POTIONS"
	case "Fer", "Cuir", "Plume divine", "Essence divine":
		return "MATÉRIAUX"
	}

	if _, existe := offgear[nom]; existe {
		return "ARMES"
	}
	if _, existe := defgear[nom]; existe {
		return "ARMURES"
	}
	return "AUTRES"
}

func (c *Character) afficherCategorieInventaire(categorie string) {
	var noms []string

	for nom, quantite := range c.Inventaire {
		if quantite > 0 && categorieObjet(nom) == categorie {
			noms = append(noms, nom)
		}
	}

	if len(noms) == 0 {
		return
	}

	sort.Strings(noms)
	sousTitre(categorie)

	for _, nom := range noms {
		fmt.Printf(
			"  %s• %-30s%s %sx%d%s\n",
			bleu, nom, reset,
			blanc, c.Inventaire[nom], reset,
		)
	}
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
	titreEcran("ARMURES")

	if len(armures) == 0 {
		fmt.Println("  Vous ne possédez aucune armure.")
	}

	for i, nom := range armures {
		armure := defgear[nom]
		fmt.Printf(
			"  %s[%d]%s %s%s%s  %s(+%d PV, %s)%s\n",
			dore, i+1, reset,
			bleu, nom, reset,
			blanc, armure.Pv, armure.slot, reset,
		)
	}

	option("R", "Retirer une armure équipée")
	option("0", "Retour")
	fmt.Print("Votre choix : ")

	var choix string
	if _, err := fmt.Scan(&choix); err != nil {
		messageErreur("Saisie invalide.")
		return
	}

	switch choix {
	case "0":
		return
	case "R", "r":
		c.armorUnequipmentMenu()
		return
	}

	var numero int
	if _, err := fmt.Sscan(choix, &numero); err != nil ||
		numero < 1 || numero > len(armures) {
		messageErreur("Choix invalide.")
		return
	}

	c.equipArmor(armures[numero-1])
}

func (c *Character) armorUnequipmentMenu() {
	titreEcran("RETIRER UNE ARMURE")
	emplacements := []string{"Casque", "Plastron", "Bottes"}

	for i, emplacement := range emplacements {
		option(
			fmt.Sprint(i+1),
			emplacement+" : "+equipmentName(c.ArmuresEquipee[emplacement]),
		)
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
	if choix < 1 || choix > len(emplacements) {
		messageErreur("Choix invalide.")
		return
	}

	c.unequipArmor(emplacements[choix-1])
}