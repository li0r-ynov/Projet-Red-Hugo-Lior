package src

import "fmt"

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
	"Glaive du Légionnaire":             {name: "Le Glaive du Légionnaire", damage: 13, vitesse: 7},
	"La Foudre de Zeus":                 {name: "La Foudre de Zeus", damage: 15, vitesse: 15},
	"Le Trident de Poséidon":            {name: "Le Trident de Poséidon", damage: 20, vitesse: 10},
	"Le Bident d'Hadès":                 {name: "le Bident d'Hadès", damage: 30, vitesse: 0},
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
		fmt.Println("Le CDP est toujours disponible.")
		return
	}

	_, exists := offgear[itemName]

	if !exists {
		fmt.Println("Cet objet n'est pas une arme.")
		return
	}

	if c.Inventaire[itemName] <= 0 {
		fmt.Println("Vous ne possédez pas cette arme.")
		return
	}

	if c.ArmeEquipee != "" {
		fmt.Printf(
			"%s est remplacé par %s.\n",
			c.ArmeEquipee,
			itemName,
		)
	}

	c.ArmeEquipee = itemName

	fmt.Printf("%s est maintenant équipé.\n", itemName)
}

func (c *Character) unequipWeapon() {
	if c.ArmeEquipee == "" {
		fmt.Println("Aucune arme n'est équipée.")
		return
	}

	fmt.Printf("%s a été retiré.\n", c.ArmeEquipee)

	c.ArmeEquipee = ""
	c.Damage = offgear["CDP"].damage
	c.updateVitesse()
}

func (c Character) cdpStats() (int, int) {
	cdp := offgear["CDP"]

	damage := c.Damage + cdp.damage
	vitesse := c.Vitesse + cdp.vitesse
	vitesse += c.armorSpeedBonus()

	return damage, vitesse
}

func (c Character) weaponStats() (int, int) {
	if c.ArmeEquipee == "" {
		return 0, 0
	}

	weapon, exists := offgear[c.ArmeEquipee]

	if !exists {
		return 0, 0
	}

	damage := c.Damage + weapon.damage
	vitesse := c.Vitesse + weapon.vitesse
	vitesse += c.armorSpeedBonus()

	return damage, vitesse
}

func (c *Character) equipArmor(itemName string) {
	armor, exists := defgear[itemName]

	if !exists {
		fmt.Println("Cet objet n'est pas une armure.")
		return
	}

	if c.Inventaire[itemName] <= 0 {
		fmt.Println("Vous ne possédez pas cette armure.")
		return
	}

	oldArmorName := c.ArmuresEquipee[armor.slot]

	if oldArmorName != "" {
		oldArmor := defgear[oldArmorName]

		c.PvMax -= oldArmor.Pv

		if c.Pv > c.PvMax {
			c.Pv = c.PvMax
		}

		fmt.Printf(
			"%s est remplacé par %s.\n",
			oldArmorName,
			itemName,
		)
	}

	c.ArmuresEquipee[armor.slot] = itemName
	c.PvMax += armor.Pv
	c.updateVitesse()

	fmt.Printf(
		"%s est équipé dans l'emplacement %s.\n",
		itemName,
		armor.slot,
	)
}

func (c *Character) updateVitesse() {
	if c.ArmeEquipee == "" {
		c.Vitesse = offgear["CDP"].vitesse
	} else {
		c.Vitesse = offgear[c.ArmeEquipee].vitesse
	}

	for _, armorName := range c.ArmuresEquipee {
		armor := defgear[armorName]
		c.Vitesse += armor.vitesse
	}
}

func (c *Character) weaponEquipmentMenu() {
	weapons := []string{}

	fmt.Println("\n=== Armes possédées ===")

	for itemName, quantity := range c.Inventaire {
		_, isWeapon := offgear[itemName]

		if isWeapon && itemName != "CDP" && quantity > 0 {
			weapons = append(weapons, itemName)

			fmt.Printf(
				"\t%d - %s\n",
				len(weapons),
				itemName,
			)
		}
	}

	if len(weapons) == 0 {
		fmt.Println("Vous ne possédez aucune arme.")
		return
	}

	fmt.Println("\t0 - Retour")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(weapons) {
		fmt.Println("Choix invalide.")
		return
	}

	weaponName := weapons[choice-1]
	c.equipWeapon(weaponName)
}
func (c Character) armorSpeedBonus() int {
	bonus := 0

	for _, armorName := range c.ArmuresEquipee {
		armor, exists := defgear[armorName]

		if exists {
			bonus += armor.vitesse
		}
	}

	return bonus
}
