package src

import (
	"fmt"
	"strings"
)

// combat renvoie true si le joueur gagne, false s'il perd.
func (c *Character) combat(adversaire opps) bool {
	resurrectionUtilisee := false
	tour := 1

	titreEcran("COMBAT")
	fmt.Printf("  %sAdversaire%s : %s%s%s\n",
		crimson, reset, bleu, adversaire.Name, reset)

	for c.Pv > 0 && adversaire.Pv > 0 {
		degatsCDP, _ := c.cdpStats()
		nomAttaqueArme := attaqueDeLArme(c.ArmeEquipee)

		titreEcran(fmt.Sprintf("TOUR %d", tour))

		fmt.Printf("  %s%-20s%s %s%d/%d PV%s\n",
			bleu, c.Name, reset, blanc, c.Pv, c.PvMax, reset)

		fmt.Printf("  %s%-20s%s %s%d/%d PV%s\n",
			bleu, adversaire.Name, reset,
			blanc, adversaire.Pv, adversaire.PvMax, reset)

		separateur()
		sousTitre("ATTAQUES")

		fmt.Printf("  %s[A]%s %sCDP%s — %s%d dégâts%s\n",
			dore, reset, bleu, reset, blanc, degatsCDP, reset)

		if nomAttaqueArme != "" {
			degatsArme, _ := c.weaponStats()
			fmt.Printf("  %s[E]%s %s%s%s — %s%d dégâts%s\n",
				dore, reset,
				bleu, nomAttaqueArme, reset,
				blanc, degatsArme, reset)
		} else {
			fmt.Printf("  %s[E]%s %sAucune arme équipée%s\n",
				dore, reset, gris, reset)
		}

		fmt.Print("Votre attaque : ")

		var choix string
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Impossible de lire votre choix.")
			return false
		}

		var degats int
		var nomAttaque string

		switch strings.ToUpper(choix) {
		case "A":
			degats = degatsCDP
			nomAttaque = "CDP"

		case "E":
			if nomAttaqueArme == "" {
				messageErreur("Équipez une arme pour utiliser cette attaque.")
				continue
			}

			degats, _ = c.weaponStats()
			nomAttaque = nomAttaqueArme

		default:
			messageErreur("Choix invalide : utilisez A ou E.")
			continue
		}

		adversaire.Pv -= degats
		if adversaire.Pv < 0 {
			adversaire.Pv = 0
		}

		fmt.Printf("\n  %s%s%s utilise %s%s%s et inflige %s%d dégâts%s !\n",
			bleu, c.Name, reset,
			crimson, nomAttaque, reset,
			blanc, degats, reset)

		if adversaire.Pv == 0 {
			break
		}

		c.dmg(adversaire.Damage)

		fmt.Printf("  %s%s%s riposte : %s%d dégâts%s reçus.\n",
			bleu, adversaire.Name, reset,
			blanc, adversaire.Damage, reset)

		if c.isDead(&resurrectionUtilisee) {
			c.Pv = c.PvMax
			fmt.Printf("  %sPV restaurés en ville : %d/%d%s\n",
				blanc, c.Pv, c.PvMax, reset)
			return false
		}

		tour++
	}

	c.Pv = c.PvMax

	titreEcran("VICTOIRE")
	messageSucces("Vous avez vaincu " + adversaire.Name + " !")
	fmt.Printf("  %sPV restaurés : %d/%d%s\n",
		blanc, c.Pv, c.PvMax, reset)

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

		sousTitre("SECONDE CHANCE")
		fmt.Printf(
			"  %sLe seigneur de cette voie vous a accordé une seconde chance...%s\n",
			blanc, reset,
		)
		fmt.Printf("  %sNe la gaspillez pas.%s\n", crimson, reset)
		fmt.Printf("  %sPV : %d/%d%s\n", blanc, c.Pv, c.PvMax, reset)
		return false
	}

	c.Pv = 0
	messageErreur("Vous êtes mort. Le combat est perdu.")
	return true
}