package src

import("fmt")



// Noms des objets de l'inventaire, utilisés comme clés dans la map.
const (
	PotionSoin   = "Potion d'ambroisie"
	PotionPoison = "Potion de poison"
	PotionPaques = "Potion de pâques"
)

// takePot consomme une potion pour soigner ou activer la potion choisie.
func (c *Character) takePot(potionChoice int) {
	var potionName string

	switch potionChoice {
	case 1:
		potionName = PotionSoin
	case 2:
		potionName = PotionPoison
	case 3:
		potionName = PotionPaques
	default:
		fmt.Println("Choix invalide.")
		return
	}

	potQuantity, potCheck := c.Inventaire[potionName]
	if !potCheck {
		fmt.Printf("Vous n'avez aucune %s dans votre inventaire.\n", potionName)
		return
	}

	if potQuantity <= 0 {
		fmt.Println("Il ne vous reste plus de potion disponible.")
		return
	}

	switch potionChoice {
	case 1:
		c.Pv += 50
		if c.Pv > c.PvMax {
			c.Pv = c.PvMax
		}
		fmt.Printf("Potion d'ambroisie utilisée ! Vous avez maintenant %d/%d PV.\n", c.Pv, c.PvMax)
	case 2:
		fmt.Println("Potion de poison utilisée, la potion va maintenant infligé 10 points de vie de dégats par seconde pendant 3 sec")
	case 3:
		fmt.Println("Vous venez d'acheter une potion de pâques ")
	}

	c.Inventaire[potionName]--
	if c.Inventaire[potionName] <= 0 {
		delete(c.Inventaire, potionName)
	}
}