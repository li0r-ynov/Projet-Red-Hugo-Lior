package src

import "fmt"

// combat renvoie true si le joueur gagne, false s'il perd.
func (c *Character) combat(adversaire opps) bool {
	resurrectionUtilisee := false

	fmt.Printf("\nUn %s apparaît ! (%d PV)\n", adversaire.Name, adversaire.Pv)

	for c.Pv > 0 && adversaire.Pv > 0 {
		fmt.Printf(
			"\n%s : %d PV | %s : %d PV\n",
			c.Name, c.Pv, adversaire.Name, adversaire.Pv,
		)

		fmt.Println("1 - Attaquer")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			adversaire.Pv -= c.Damage

			if adversaire.Pv < 0 {
				adversaire.Pv = 0
			}

			fmt.Printf(
				"Vous infligez %d dégâts. Il reste %d PV à %s.\n",
				c.Damage, adversaire.Pv, adversaire.Name,
			)

		default:
			fmt.Println("Choix invalide : vous ne perdez pas votre tour.")
			continue
		}

		// L'adversaire ne joue pas s'il vient d'être vaincu.
		if adversaire.Pv == 0 {
			break
		}

		c.dmg(adversaire.Damage)
		fmt.Printf(
			"%s vous inflige %d dégâts.\n",
			adversaire.Name, adversaire.Damage,
		)

		if c.isDead(&resurrectionUtilisee) {
			return false
		}
	}

	fmt.Printf("Vous avez vaincu %s !\n", adversaire.Name)
	return true
}

func (c *Character) dmg(damage int) {
	c.Pv -= damage

	if c.Pv < 0 {
		c.Pv = 0
	}
}

// isDead renvoie true uniquement si le joueur meurt après
// avoir déjà utilisé sa résurrection pendant ce combat.
func (c *Character) isDead(resurrectionUtilisee *bool) bool {
	if c.Pv > 0 {
		return false
	}

	if !*resurrectionUtilisee {
		*resurrectionUtilisee = true
		c.Pv = c.PvMax / 2

		fmt.Println("Le seigneur de cette voie vous a accordé une seconde chance... Ne la gaspillez pas.")
		fmt.Printf("Vous ressuscitez avec %d PV.\n", c.Pv)
		return false
	}

	c.Pv = 0
	fmt.Println("Vous êtes mort.")
	return true
}