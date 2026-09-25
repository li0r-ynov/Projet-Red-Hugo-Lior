package src

import (
	"fmt"
	"sort"
	"strings"
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
	"Casque de Gladiateur":   {name: "Casque de Gladiateur", Pv: 10, slot: "Casque"},
	"Plastron de Gladiateur": {name: "Plastron de Gladiateur", Pv: 15, slot: "Plastron"},
	"Bottes de Gladiateur":   {name: "Bottes de Gladiateur", Pv: 5, slot: "Bottes", vitesse: 5},
	"Bottes de Hermès":       {name: "Bottes de Hermès", Pv: 5, slot: "Bottes", vitesse: 15},
	"Ailes d'Icare":          {name: "Ailes d'Icare", Pv: 10, slot: "Plastron", vitesse: 10},
	"Casque d'Arès":          {name: "Casque d'Arès", Pv: 20, slot: "Casque"},
}

func (c *Character) equipWeapon(nom string) {
	if _, existe := offgear[nom]; !existe || nom == "CDP" {
		messageErreur("Arme inconnue.")
		return
	}
	if c.Inventaire[nom] <= 0 {
		messageErreur("Vous ne possédez pas cette arme.")
		return
	}
	if c.ArmeEquipee == nom {
		messageErreur("Cette arme est déjà équipée.")
		return
	}

	c.ArmeEquipee = nom
	messageSucces(nom + " est équipé.")
}

func (c *Character) unequipWeapon() {
	if c.ArmeEquipee == "" {
		messageErreur("Aucune arme équipée.")
		return
	}

	messageSucces(c.ArmeEquipee + " a été retirée.")
	c.ArmeEquipee = ""
}

func (c Character) armorSpeedBonus() int {
	bonus := 0
	for _, nom := range c.ArmuresEquipee {
		bonus += defgear[nom].vitesse
	}
	return bonus
}

func (c Character) cdpStats() (int, int) {
	cdp := offgear["CDP"]
	return c.Damage + cdp.damage,
		c.Vitesse + cdp.vitesse + c.armorSpeedBonus()
}

func (c Character) weaponStats() (int, int) {
	arme, existe := offgear[c.ArmeEquipee]
	if !existe || c.ArmeEquipee == "" {
		return 0, 0
	}

	return c.Damage + arme.damage,
		c.Vitesse + arme.vitesse + c.armorSpeedBonus()
}

func (c *Character) weaponEquipmentMenu() {
	var armes []string

	for nom, quantite := range c.Inventaire {
		_, existe := offgear[nom]
		if existe && nom != "CDP" && quantite > 0 {
			armes = append(armes, nom)
		}
	}

	sort.Strings(armes)
	titreEcran("ARMES")

	if len(armes) == 0 {
		fmt.Println("  Vous ne possédez aucune arme.")
	}

	for i, nom := range armes {
		fmt.Printf(
			"  %s[%d]%s %s%s%s  %s(+%d dégâts)%s\n",
			dore, i+1, reset,
			bleu, nom, reset,
			blanc, offgear[nom].damage, reset,
		)
	}

	option("R", "Retirer l'arme équipée : "+equipmentName(c.ArmeEquipee))
	option("0", "Retour")
	fmt.Print("Votre choix : ")

	var choix string
	if _, err := fmt.Scan(&choix); err != nil {
		messageErreur("Saisie invalide.")
		return
	}

	if strings.EqualFold(choix, "R") {
		c.unequipWeapon()
		return
	}
	if choix == "0" {
		return
	}

	var numero int
	if _, err := fmt.Sscan(choix, &numero); err != nil ||
		numero < 1 || numero > len(armes) {
		messageErreur("Choix invalide.")
		return
	}

	c.equipWeapon(armes[numero-1])
}

func (c *Character) equipArmor(nom string) {
	armure, existe := defgear[nom]
	if !existe {
		messageErreur("Armure inconnue.")
		return
	}
	if c.Inventaire[nom] <= 0 {
		messageErreur("Vous ne possédez pas cette armure.")
		return
	}
	if c.ArmuresEquipee == nil {
		c.ArmuresEquipee = make(map[string]string)
	}
	if c.ArmuresEquipee[armure.slot] == nom {
		messageErreur("Cette armure est déjà équipée.")
		return
	}

	ancienne := c.ArmuresEquipee[armure.slot]
	if ancienne != "" {
		c.PvMax -= defgear[ancienne].Pv
		if c.Pv > c.PvMax {
			c.Pv = c.PvMax
		}
	}

	c.ArmuresEquipee[armure.slot] = nom
	c.PvMax += armure.Pv
	messageSucces(nom + " est équipé.")
}

func (c *Character) unequipArmor(emplacement string) {
	nom := c.ArmuresEquipee[emplacement]
	if nom == "" {
		messageErreur("Aucune armure équipée dans cet emplacement.")
		return
	}

	c.PvMax -= defgear[nom].Pv
	if c.Pv > c.PvMax {
		c.Pv = c.PvMax
	}

	delete(c.ArmuresEquipee, emplacement)
	messageSucces(nom + " a été retiré.")
}