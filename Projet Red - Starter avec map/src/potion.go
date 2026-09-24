package src

import "fmt"

const (
	PotionSoin = "Potion d'ambroisie"
	PotionPoison = "Potion de poison"
	PotionPaque = "Potion de Pâques"
)

func (c *Character) takePot(choix int) {
	var nom string
	switch choix {
	case 1: nom = PotionSoin
	case 2: nom = PotionPoison
	case 3:
		messageErreur("La potion de Pâques s'utilise dans une Maison des Dieux.")
		return
	default: messageErreur("Choix invalide."); return
	}
	if c.Inventaire[nom] == 0 { messageErreur("Vous ne possédez pas cet objet."); return }
	switch choix {
	case 1:
		if c.Pv >= c.PvMax { messageErreur("Vous avez déjà tous vos PV."); return }
		c.Pv += 30
		if c.Pv > c.PvMax { c.Pv = c.PvMax }
		c.removeInventory(nom, 1)
		messageSucces(fmt.Sprintf("Ambroisie utilisée : %d/%d PV.", c.Pv, c.PvMax))
	case 2:
		// L'effet sur un ennemi doit être défini avec le système de combat.
		messageErreur("Le poison ne peut pas encore être utilisé hors combat.")
	}
}
