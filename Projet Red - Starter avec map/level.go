package main

import "fmt"

func (c *Character) ChangeRenown(nombre int) {
	c.Renommee += nombre
	c.updateLevel()
}

func (c *Character) updateLevel() {
	switch {
	case c.Renommee >= 3100:
		c.Level = "Dieu de l'Olympe"

	case c.Renommee >= 1500:
		c.Level = "Demi-dieu"

	case c.Renommee >= 700:
		c.Level = "Héros"

	case c.Renommee >= 300:
		c.Level = "Guerrier"

	case c.Renommee >= 100:
		c.Level = "Combattant"

	case c.Renommee <= -3100:
		c.Level = "Dieu des Enfers"

	case c.Renommee <= -1500:
		c.Level = "Demi-dieu"

	case c.Renommee <= -700:
		c.Level = "Démon"

	case c.Renommee <= -300:
		c.Level = "Mercenaire"

	case c.Renommee <= -100:
		c.Level = "Bandit"

	default:
		c.Level = "Civil"
	}
}

	func (c *Character) Title() string {
	titles := map[int]string{
		5:  "Dieu de l'Olympe",
		4:  "Demi-dieu",
		3:  "Héros",
		2:  "Guerrier",
		1:  "Combattant",
		0:  "Civil",
		-1: "Bandit",
		-2: "Mercenaire",
		-3: "Démon",
		-4: "Demi-dieu",
		-5: "Dieu des Enfers",
	}

	return titles[c.Level]
}