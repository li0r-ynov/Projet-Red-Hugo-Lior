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
	money                   int
	PotionGratuiteRecuperee bool
	LimitInventaire         int
}

// initCharacter initialise un personnage selon sa classe (PV max différents)
// et lui donne un inventaire de départ.
func (c *Character) initCharacter(name string, class string) {
	c.Name = name
	c.Classe = class
	c.LimitInventaire = 10
	c.money = 10
	switch c.Classe {
	case "athénien ":
		c.PvMax = 100
		c.Pv = c.PvMax / 2
	case "civil":
		c.PvMax = 100
		c.Pv = c.PvMax / 2
	}
	c.Inventaire = map[string]int{
		PotionSoin:   1,
		PotionPoison: 1,
		PotionPaques: 1,
	}
}

// displaylnfo affiche les informations principales du personnage.
func (c Character) displaylnfo() {
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
	fmt.Printf("\t Pièces d'or : %d\n", c.money)

}
