package src

import (
	"fmt"
	"sort"
)

type Weapon struct {
	name    string
	damage  int
	vitesse int
}

var offgear = map[string]Weapon{
	"CDP":                               {name: "CDP", damage: 5, vitesse: 6},
	"Les Dagues de l'Assassin":          {name: "Les Dagues de l'Assassin", damage: 10, vitesse: 10},
	"La Hache de Viking":                {name: "La Hache de Viking", damage: 15, vitesse: 5},
	"Le Marteau de Guerre du Martelier": {name: "Le Marteau de Guerre du Martelier", damage: 17, vitesse: 0},
	"Le Glaive du Légionnaire":          {name: "Le Glaive du Légionnaire", damage: 13, vitesse: 7},
	"La Foudre de Zeus":                 {name: "La Foudre de Zeus", damage: 15, vitesse: 15},
	"Le Trident de Poséidon":            {name: "Le Trident de Poséidon", damage: 20, vitesse: 10},
	"Le Bident d'Hadès":                 {name: "Le Bident d'Hadès", damage: 30, vitesse: 0},
}

type Armor struct {
	name    string
	Pv      int
	slot    string
	vitesse int
}

var defgear = map[string]Armor{
	"Casque de Gladiateur":   {name: "Casque de Gladiateur", Pv: 10, slot: "Casque", vitesse: 0},
	"Plastron de Gladiateur": {name: "Plastron de Gladiateur", Pv: 15, slot: "Plastron", vitesse: 0},
	"Bottes de Gladiateur":   {name: "Bottes de Gladiateur", Pv: 5, slot: "Bottes", vitesse: 5},
	"Bottes de Hermès":       {name: "Bottes de Hermès", Pv: 5, slot: "Bottes", vitesse: 15},
	"Ailes d'Icare":          {name: "Ailes d'Icare", Pv: 10, slot: "Plastron", vitesse: 10},
	"Casque d'Arès":          {name: "Casque d'Arès", Pv: 20, slot: "Casque", vitesse: 0},
}

func (c *Character) equipWeapon(itemName string) {
	if itemName == "CDP" {
		messageSucces("Le CDP est toujours disponible.")
		return
	}

	if _, existe := offgear[itemName]; !existe {
		messageErreur("Cet objet n'est pas une arme.")
		return
	}

	if c.Inventaire[itemName] <= 0 {
		messageErreur("Vous ne possédez pas cette arme.")
		return
	}

	if c.ArmeEquipee == itemName {
		messageSucces("Cette arme est déjà équipée.")
		return
	}

	if c.ArmeEquipee != "" {
		fmt.Printf("  %s%s%s est remplacé par %s%s%s.\n",
			bleu, c.ArmeEquipee, reset,
			bleu, itemName, reset)
	}

	c.ArmeEquipee = itemName
	c.updateVitesse()
	messageSucces(itemName + " est maintenant équipé.")
}

func (c *Character) unequipWeapon() {
	if c.ArmeEquipee == "" {
		messageErreur("Aucune arme n'est équipée.")
		return
	}

	ancienneArme := c.ArmeEquipee
	c.ArmeEquipee = ""
	c.updateVitesse()
	messageSucces(ancienneArme + " a été retirée.")
}

func (c Character) cdpStats() (int, int) {
	cdp := offgear["CDP"]

	degats := c.Damage + cdp.damage
	vitesse := c.Vitesse + cdp.vitesse + c.armorSpeedBonus()

	return degats, vitesse
}

func (c Character) weaponStats() (int, int) {
	if c.ArmeEquipee == "" {
		return 0, 0
	}

	arme, existe := offgear[c.ArmeEquipee]
	if !existe {
		return 0, 0
	}

	degats := c.Damage + arme.damage
	vitesse := c.Vitesse + arme.vitesse + c.armorSpeedBonus()

	return degats, vitesse
}

func (c *Character) equipArmor(itemName string) {
	armure, existe := defgear[itemName]
	if !existe {
		messageErreur("Cet objet n'est pas une armure.")
		return
	}

	if c.Inventaire[itemName] <= 0 {
		messageErreur("Vous ne possédez pas cette armure.")
		return
	}

	if c.ArmuresEquipee == nil {
		c.ArmuresEquipee = make(map[string]string)
	}

	ancienne := c.ArmuresEquipee[armure.slot]
	if ancienne == itemName {
		messageSucces(itemName + " est déjà équipé.")
		return
	}

	if ancienne != "" {
		c.PvMax -= defgear[ancienne].Pv

		if c.Pv > c.PvMax {
			c.Pv = c.PvMax
		}

		fmt.Printf("  %s%s%s est remplacé par %s%s%s.\n",
			bleu, ancienne, reset,
			bleu, itemName, reset)
	}

	c.ArmuresEquipee[armure.slot] = itemName
	c.PvMax += armure.Pv
	c.updateVitesse()

	switch armure.slot {
	case "Bottes":
		messageSucces(itemName + " sont équipées.")
	case "Plastron":
		if itemName == "Ailes d'Icare" {
			messageSucces(itemName + " sont équipées.")
		} else {
			messageSucces(itemName + " est équipé.")
		}
	default:
		messageSucces(itemName + " est équipé.")
	}
}

func (c *Character) updateVitesse() {
	// Vitesse conserve ici la valeur naturelle de la classe.
	// Le bonus de l'arme et des armures est ajouté par les fonctions de stats.
	c.Vitesse = 10
}

func (c *Character) weaponEquipmentMenu() {
	var armes []string

	for nom, quantite := range c.Inventaire {
		_, estUneArme := offgear[nom]
		if estUneArme && nom != "CDP" && quantite > 0 {
			armes = append(armes, nom)
		}
	}

	sort.Strings(armes)
	titreEcran("ÉQUIPER UNE ARME")

	if len(armes) == 0 {
		fmt.Println("  Vous ne possédez aucune arme.")
		return
	}

	for i, nom := range armes {
		statistiques := offgear[nom]
		fmt.Printf("  %s[%d]%s %s%s%s  %s(+%d dégâts)%s\n",
			dore, i+1, reset,
			bleu, nom, reset,
			blanc, statistiques.damage, reset)
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

	if choix < 1 || choix > len(armes) {
		messageErreur("Choix invalide.")
		return
	}

	c.equipWeapon(armes[choix-1])
}

func (c Character) armorSpeedBonus() int {
	bonus := 0

	for _, nom := range c.ArmuresEquipee {
		if armure, existe := defgear[nom]; existe {
			bonus += armure.vitesse
		}
	}

	return bonus
}