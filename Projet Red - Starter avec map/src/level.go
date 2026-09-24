package src

import "fmt"

func (c *Character) ChangeRenown(nombre int) {
	c.Renown += nombre
	c.updateLevel()
}

func (c *Character) updateLevel() {
	ancienTitre := c.Level
	palierActuel := 0

	switch {
	case c.Renown >= 3100:
		c.Level = "Dieu de l'Olympe"
		palierActuel = 5

	case c.Renown >= 1500:
		c.Level = "Demi-dieu"
		palierActuel = 4

	case c.Renown >= 700:
		c.Level = "Héros"
		palierActuel = 3

	case c.Renown >= 300:
		c.Level = "Guerrier"
		palierActuel = 2

	case c.Renown >= 100:
		c.Level = "Combattant"
		palierActuel = 1

	case c.Renown <= -3100:
		c.Level = "Dieu des Enfers"
		palierActuel = 5

	case c.Renown <= -1500:
		c.Level = "Demi-dieu"
		palierActuel = 4

	case c.Renown <= -700:
		c.Level = "Démon"
		palierActuel = 3

	case c.Renown <= -300:
		c.Level = "Mercenaire"
		palierActuel = 2

	case c.Renown <= -100:
		c.Level = "Bandit"
		palierActuel = 1

	default:
		c.Level = "Civil"
	}

	if palierActuel > c.PalierPvAtteint {
		paliersGagnes := palierActuel - c.PalierPvAtteint
		pvGagnes := paliersGagnes * 10

		c.PvMax += pvGagnes
		c.Pv += pvGagnes
		c.PalierPvAtteint = palierActuel

		fmt.Printf(
			"Niveau supérieur ! +%d PV. Vous avez %d/%d PV.\n",
			pvGagnes, c.Pv, c.PvMax,
		)
	}

	if c.Level != ancienTitre {
		fmt.Printf("Nouveau titre : %s !\n", c.Level)
	}
}