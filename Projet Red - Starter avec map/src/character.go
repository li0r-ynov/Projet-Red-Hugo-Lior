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
}

// initCharacter initialise un personnage selon sa classe (PV max différents)
// et lui donne un inventaire de départ.
func (c *Character) initCharacter() {
	c.LimitInventaire = 10
	switch c.Classe {
	case "Athénien":
		c.PvMax = 60
		c.Pv = c.PvMax / 2
		c.Money = 10
		c.Level = "Civil"
		c.Renown = 0
	case "Sparte":
		c.PvMax = 50
		c.Pv = c.PvMax / 2
		c.Money = 10
		c.Level = "Civil"
		c.Renown = 0
	}
	c.Inventaire = map[string]int{
		PotionSoin:   0,
		PotionPoison: 0,
		PotionPaques: 0,
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
	fmt.Printf("\t Pièce d'or : %d\n", c.Money)
	fmt.Printf("Titre : %s\n", c.Level)
	fmt.Printf("Renommée : %d\n", c.Renown)
}
