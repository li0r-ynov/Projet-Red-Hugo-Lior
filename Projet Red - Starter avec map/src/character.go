package src

import (
	"fmt"
)

type Character struct {
	Name                    string
	Classe                  string
	PvMax                   int
	Pv                      int
	Inventaire              map[string]int
	Money                   int
	Level                   string
	Renown                  int
	PotionGratuiteRecuperee bool
	LimitInventaire         int
	UpgradeCount            int
	Damage                  int
	Vitesse                 int
	ArmeEquipee             string
	ArmuresEquipee          map[string]string
<<<<<<< HEAD
=======
	DirectionTour           string
	EtageTour               int
>>>>>>> 67fbb299963fb932cd9c41a83a5c0d3c2a6cf7e0
}

// initCharacter initialise un personnage selon sa classe (PV max différents)
// et lui donne un inventaire de départ.
func (c *Character) initCharacter() {
	c.LimitInventaire = 10

	switch c.Classe {
	case "Athénien":
		c.PvMax = 50
		c.Pv = c.PvMax / 2
		c.Money = 100
		c.Damage = 15
		c.Vitesse = 10

	case "Sparte":
		c.PvMax = 60
		c.Pv = c.PvMax / 2
		c.Money = 100
		c.Damage = 10
		c.Vitesse = 10
	}

	c.Level = "Civil"
	c.Renown = 0

	c.Inventaire = make(map[string]int)
	c.ArmuresEquipee = make(map[string]string)
	c.ArmeEquipee = ""
	c.DirectionTour = ""
	c.EtageTour = 1
}

// displaylnfo affiche les informations principales du personnage.
func (c Character) displayInfo() {
	totalItems := 0
	for _, quantity := range c.Inventaire {
		totalItems += quantity
	}

	fmt.Println("=== Fiche du personnage ===")
	fmt.Printf("\t Nom      : %s\n", c.Name)
	fmt.Printf("\t Classe   : %s\n", c.Classe)
	fmt.Printf("\t PV       : %d\n", c.Pv)
	fmt.Printf("\t PV max   : %d\n", c.PvMax)
	fmt.Printf("\t Inventaire : %d objet(s)\n", totalItems)
	fmt.Printf("\t Pièce d'or : %d\n", c.Money)
	fmt.Printf("Titre : %s\n", c.Level)
	fmt.Printf("Renommée : %d\n", c.Renown)

	fmt.Println("=== Équipement ===")

	if c.ArmeEquipee == "" {
		fmt.Println("\t Arme équipée : Aucune")
		fmt.Println("\t Attaque disponible : CDP")
	} else {
		fmt.Printf("\t Arme équipée : %s\n", c.ArmeEquipee)
		fmt.Println("\t Attaque disponible : CDP")
	}

	fmt.Printf(
		"\t Casque : %s\n",
		equipmentName(c.ArmuresEquipee["Casque"]),
	)

	fmt.Printf(
		"\t Plastron : %s\n",
		equipmentName(c.ArmuresEquipee["Plastron"]),
	)

	fmt.Printf(
		"\t Bottes : %s\n",
		equipmentName(c.ArmuresEquipee["Bottes"]),
	)

	fmt.Println("=== Statistiques de combat ===")

	fmt.Printf("\t Dégâts naturels   : %d\n", c.Damage)
	fmt.Printf("\t Vitesse naturelle : %d\n", c.Vitesse)

	cdpDamage, cdpVitesse := c.cdpStats()

	fmt.Println("\t CDP :")
	fmt.Printf("\t   Dégâts  : %d\n", cdpDamage)
	fmt.Printf("\t   Vitesse : %d\n", cdpVitesse)

	switch c.ArmeEquipee {
	case "":
		fmt.Println("\t Arme équipée : Aucune")

	default:
		weaponDamage, weaponVitesse := c.weaponStats()

		fmt.Printf("\t %s :\n", c.ArmeEquipee)
		fmt.Printf("\t   Dégâts  : %d\n", weaponDamage)
		fmt.Printf("\t   Vitesse : %d\n", weaponVitesse)
	}
}

func equipmentName(itemName string) string {
	if itemName == "" {
		return "Vide"
	}

	return itemName
}
<<<<<<< HEAD

func (c *Character) UpgradeInventorySlot() bool {
	if c.UpgradeCount < 3 {
		c.LimitInventaire += 10
		c.UpgradeCount++
		return true
	}
	return false
}
=======
>>>>>>> 67fbb299963fb932cd9c41a83a5c0d3c2a6cf7e0
