package src



func (c *Character) ChangeRenown(nombre int) {
	c.Renown += nombre
	c.updateLevel()
}

func (c *Character) updateLevel() {
	switch {
	case c.Renown >= 3100:
		c.Level = "Dieu de l'Olympe"

	case c.Renown >= 1500:
		c.Level = "Demi-dieu"

	case c.Renown >= 700:
		c.Level = "Héros"

	case c.Renown >= 300:
		c.Level = "Guerrier"

	case c.Renown >= 100:
		c.Level = "Combattant"

	case c.Renown <= -3100:
		c.Level = "Dieu des Enfers"

	case c.Renown <= -1500:
		c.Level = "Demi-dieu"

	case c.Renown <= -700:
		c.Level = "Démon"

	case c.Renown <= -300:
		c.Level = "Mercenaire"

	case c.Renown <= -100:
		c.Level = "Bandit"

	default:
		c.Level = "Civil"
	}
}
