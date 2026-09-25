package src

import "fmt"

func (c *Character) ChangeRenown(nombre int) {
	c.Renown += nombre
	c.updateLevel()
}

func (c *Character) updateLevel() {
	ancienTitre := c.Level
	palier := 0

	switch {
	case c.Renown >= 3100:
		c.Level, palier = "Dieu de l'Olympe", 5
	case c.Renown >= 1500:
		c.Level, palier = "Demi-dieu", 4
	case c.Renown >= 700:
		c.Level, palier = "Héros", 3
	case c.Renown >= 300:
		c.Level, palier = "Guerrier", 2
	case c.Renown >= 100:
		c.Level, palier = "Combattant", 1
	case c.Renown <= -3100:
		c.Level, palier = "Dieu des Enfers", 5
	case c.Renown <= -1500:
		c.Level, palier = "Demi-dieu", 4
	case c.Renown <= -700:
		c.Level, palier = "Démon", 3
	case c.Renown <= -300:
		c.Level, palier = "Mercenaire", 2
	case c.Renown <= -100:
		c.Level, palier = "Bandit", 1
	default:
		c.Level = "Civil"
	}

	if palier > c.PalierPvAtteint {
		gain := (palier - c.PalierPvAtteint) * 10
		c.PvMax += gain
		c.Pv += gain
		c.PalierPvAtteint = palier

		messageSucces(fmt.Sprintf(
			"Niveau supérieur : +%d PV (%d/%d).",
			gain, c.Pv, c.PvMax,
		))
	}

	if c.Level != ancienTitre {
		messageSucces("Nouveau titre : " + c.Level)
	}
}