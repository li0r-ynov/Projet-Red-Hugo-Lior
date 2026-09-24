package src

import (
	"fmt"
)

// Character représente un personnage jouable avec ses points de vie et son inventaire.
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
	Damage                  int
	Vitesse                 int
}

// initCharacter initialise un personnage selon sa classe (PV max différents)
// et lui donne un inventaire de départ.
func (c *Character) initCharacter() {
	switch c.Classe {
	case "Athénien":
		c.PvMax = 60
		c.Pv = c.PvMax / 2
		c.Money = 100
		c.Level = "Civil"
		c.Renown = 0
		c.LimitInventaire = 10
	case "Sparte":
		c.PvMax = 50
		c.Pv = c.PvMax / 2
		c.Money = 100
		c.Level = "Civil"
		c.Renown = 0
		c.LimitInventaire = 10
	}
	c.Inventaire = map[string]int{
		PotionSoin:   0,
		PotionPoison: 0,
		PotionPaque:  0,
	}

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
}

func (c *Character) isDead() bool {
	if c.Pv <= 0 {
		return true
	}
	return false
}
