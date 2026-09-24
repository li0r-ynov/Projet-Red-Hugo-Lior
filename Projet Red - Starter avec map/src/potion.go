package src

import (
	"fmt"
	"time"
)

const (
	PotionSoin   = "Potion d'ambroisie"
	PotionPoison = "Potion de poison"
	PotionPaque  = "Potion de Pâques"
)

// poisonPot applique trois fois 10 dégâts, à une seconde d'intervalle.
func poisonPot(adversaire *opps) {
	for seconde := 1; seconde <= 3; seconde++ {
		time.Sleep(1 * time.Second)

		adversaire.Pv -= 10
		if adversaire.Pv < 0 {
			adversaire.Pv = 0
		}

		fmt.Printf(
			"  %sPoison • seconde %d%s : %s%d/%d PV%s pour %s%s%s\n",
			crimson, seconde, reset,
			blanc, adversaire.Pv, adversaire.PvMax, reset,
			bleu, adversaire.Name, reset,
		)

		if adversaire.Pv == 0 {
			return
		}
	}
}

func (c *Character) takePot(choix int) {
	switch choix {
	case 1:
		if c.Inventaire[PotionSoin] <= 0 {
			messageErreur("Vous ne possédez pas de potion d'ambroisie.")
			return
		}

		if c.Pv >= c.PvMax {
			messageErreur("Vous avez déjà tous vos PV.")
			return
		}

		c.Pv += 30
		if c.Pv > c.PvMax {
			c.Pv = c.PvMax
		}

		c.removeInventory(PotionSoin, 1)
		messageSucces(fmt.Sprintf(
			"Ambroisie utilisée : %d/%d PV.",
			c.Pv, c.PvMax,
		))

	case 2:
		messageErreur("La potion de poison s'utilise pendant un combat avec la touche P.")

	case 3:
		messageErreur("La potion de Pâques s'utilise dans une Maison des Dieux.")

	default:
		messageErreur("Choix invalide.")
	}
}