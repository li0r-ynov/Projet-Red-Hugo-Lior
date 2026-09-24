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

		fmt.Println("1 - CDP")

		nomAttaque := attaqueDeLArme(c.ArmeEquipee)
		if nomAttaque != "" {
			fmt.Printf("2 - %s (%s)\n", nomAttaque, c.ArmeEquipee)
		}

		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("Saisie invalide.")
			return false
		}

		var degats int
		var attaque string

		switch choix {
		case 1:
			degats, _ = c.cdpStats()
			attaque = "CDP"

		case 2:
			if nomAttaque == "" {
				fmt.Println("Vous n'avez pas d'attaque d'arme disponible.")
				continue
			}

			degats, _ = c.weaponStats()
			attaque = nomAttaque

		default:
			fmt.Println("Choix invalide : vous ne perdez pas votre tour.")
			continue
		}

		adversaire.Pv -= degats
		if adversaire.Pv < 0 {
			adversaire.Pv = 0
		}

		fmt.Printf(
			"%s inflige %d dégâts. Il reste %d PV à %s.\n",
			attaque, degats, adversaire.Pv, adversaire.Name,
		)

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

func attaqueDeLArme(arme string) string {
	switch arme {
	case "Le Marteau de Guerre du Martelier":
		return "Le coup du marteau"

	case "Les Dagues de l'Assassin":
		return "Entaillade"

	case "Le Glaive du Légionnaire":
		return "Lame sanglante"

	case "La Hache de Viking":
		return "Frappe de barbare"

	default:
		return ""
	}
}

func (c *Character) dmg(damage int) {
	c.Pv -= damage

	if c.Pv < 0 {
		c.Pv = 0
	}
}

// isDead renvoie true si la seconde chance a déjà été utilisée.
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