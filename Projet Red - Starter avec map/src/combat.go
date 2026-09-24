package src

import (
	"fmt"
	"strings"
)

func attaqueDeLArme(nom string) string {
	switch nom {
	case "Les Dagues de l'Assassin": return "Entaillade"
	case "La Hache de Viking": return "Frappe de barbare"
	case "Le Marteau de Guerre du Martelier": return "Le coup du marteau"
	case "Le Glaive du Légionnaire": return "Lame sanglante"
	case "La Foudre de Zeus": return "Foudre divine"
	case "Le Trident de Poséidon": return "Tsunami"
	case "Le Bident d'Hadès": return "Feu des Enfers"
	default: return ""
	}
}

// combat retourne true en cas de victoire. Les PV sont restaurés à la fin.
func (c *Character) combat(adversaire opps) bool {
	resurrectionUtilisee := false
	tour := 1
	titreEcran("COMBAT CONTRE " + adversaire.Name)
	for c.Pv > 0 && adversaire.Pv > 0 {
		titreEcran(fmt.Sprintf("TOUR %d", tour))
		ligneInfo(c.Name, fmt.Sprintf("%d/%d PV", c.Pv, c.PvMax))
		ligneInfo(adversaire.Name, fmt.Sprintf("%d/%d PV", adversaire.Pv, adversaire.PvMax))
		separateur()
		sousTitre("ATTAQUES")
		degatsCDP, _ := c.cdpStats()
		fmt.Printf("  %s[A]%s %sCDP%s — %s%d dégâts%s\n", dore, reset, bleu, reset, blanc, degatsCDP, reset)
		nomArme := attaqueDeLArme(c.ArmeEquipee)
		if nomArme != "" {
			degatsArme, _ := c.weaponStats()
			fmt.Printf("  %s[E]%s %s%s%s — %s%d dégâts%s\n", dore, reset, bleu, nomArme, reset, blanc, degatsArme, reset)
		} else { fmt.Printf("  %s[E]%s %sAucune arme équipée%s\n", dore, reset, gris, reset) }
		fmt.Print("Votre attaque : ")
		var choix string
		if _, err := fmt.Scan(&choix); err != nil { messageErreur("Saisie invalide."); return false }
		degats := 0
		attaque := ""
		switch strings.ToUpper(choix) {
		case "A": degats, attaque = degatsCDP, "CDP"
		case "E":
			if nomArme == "" { messageErreur("Équipez une arme avant d'utiliser E."); continue }
			degats, _ = c.weaponStats()
			attaque = nomArme
		default: messageErreur("Utilisez A ou E."); continue
		}
		adversaire.Pv -= degats
		if adversaire.Pv < 0 { adversaire.Pv = 0 }
		messageSucces(fmt.Sprintf("%s utilise %s : %d dégâts.", c.Name, attaque, degats))
		if adversaire.Pv == 0 { break }
		c.dmg(adversaire.Damage)
		messageErreur(fmt.Sprintf("%s inflige %d dégâts.", adversaire.Name, adversaire.Damage))
		if c.isDead(&resurrectionUtilisee) {
			c.Pv = c.PvMax
			ligneInfo("Retour en ville", fmt.Sprintf("%d/%d PV", c.Pv, c.PvMax))
			return false
		}
		tour++
	}
	c.Pv = c.PvMax
	titreEcran("VICTOIRE")
	messageSucces("Vous avez vaincu " + adversaire.Name + " !")
	ligneInfo("PV restaurés", fmt.Sprintf("%d/%d", c.Pv, c.PvMax))
	return true
}

func (c *Character) dmg(degats int) {
	c.Pv -= degats
	if c.Pv < 0 { c.Pv = 0 }
}

// isDead retourne true à la deuxième mort dans ce combat.
func (c *Character) isDead(utilisee *bool) bool {
	if c.Pv > 0 { return false }
	if !*utilisee {
		*utilisee = true
		c.Pv = c.PvMax / 2
		sousTitre("SECONDE CHANCE")
		fmt.Printf("  %sLe seigneur de cette voie vous a accordé une seconde chance... Ne la gaspillez pas.%s\n", blanc, reset)
		ligneInfo("PV", fmt.Sprintf("%d/%d", c.Pv, c.PvMax))
		return false
	}
	c.Pv = 0
	messageErreur("Vous êtes mort. Le combat est perdu.")
	return true
}
