package src

import "fmt"

func (c *Character) TowerTravelDisplay(text, a, b, d, e, retour string) int {
	fmt.Println(text)
	option("1", a); option("2", b); option("3", d); option("4", e); option("0", retour)
	fmt.Print("Votre choix : ")
	var choix int
	fmt.Scan(&choix)
	return choix
}

func (c *Character) MenuTour() {
	for c.DirectionTour == "" {
		titreEcran("LA TOUR")
		fmt.Printf("  %sVotre choix sera définitif.%s\n", crimson, reset)
		option("1", "Monter vers l'Olympe")
		option("2", "Descendre vers les Enfers")
		option("0", "Retour à Athènes")
		fmt.Print("Votre choix : ")
		var choix int
		if _, err := fmt.Scan(&choix); err != nil { messageErreur("Saisie invalide."); return }
		switch choix {
		case 0: return
		case 1: c.DirectionTour = "Monter"
		case 2: c.DirectionTour = "Descendre"
		default: messageErreur("Choix invalide.")
		}
	}
	for {
		if c.EtageTour > 5 { titreEcran("VOIE TERMINÉE"); messageSucces("Cinq combats remportés."); return }
		nom := adversaireEtage(c.DirectionTour, c.EtageTour)
		adversaire, ok := oppsdef[nom]
		if !ok { messageErreur("Adversaire introuvable : " + nom); return }
		titreEcran("PARCOURS DE LA TOUR")
		ligneInfo("Voie", c.DirectionTour)
		ligneInfo("Étage", fmt.Sprintf("%d/5", c.EtageTour))
		ligneInfo("Adversaire", adversaire.Name)
		option("1", "Commencer le combat")
		option("0", "Retour à Athènes")
		fmt.Print("Votre choix : ")
		var choix int
		if _, err := fmt.Scan(&choix); err != nil { messageErreur("Saisie invalide."); return }
		switch choix {
		case 0: return
		case 1:
			if c.Pv <= 0 { messageErreur("Récupérez des PV avant de combattre."); return }
			etage := c.EtageTour
			if !c.combat(adversaire) { messageErreur("Vous pourrez retenter cet étage."); return }
			c.recompenseEtage(etage)
			c.EtageTour++
			messageSucces("Continuez ou retournez en ville.")
		default: messageErreur("Choix invalide.")
		}
	}
}

func adversaireEtage(direction string, etage int) string {
	monter := []string{"", "Combattant", "Guerrier", "Héro", "Demi-Dieu", "Zeus"}
	descendre := []string{"", "Bandit", "Mercenaire", "Démon", "Demi-Dieu", "Hadès"}
	if etage < 1 || etage > 5 { return "" }
	if direction == "Monter" { return monter[etage] }
	if direction == "Descendre" { return descendre[etage] }
	return ""
}

func (c *Character) recompenseEtage(etage int) {
	seuils := []int{0, 100, 300, 700, 1500, 3100}
	renown := seuils[etage]
	if c.DirectionTour == "Descendre" { renown = -renown }
	c.ChangeRenown(renown - c.Renown)
	sousTitre("RÉCOMPENSES")
	ligneInfo("Titre", c.Level)
	ligneInfo("Renommée", fmt.Sprint(c.Renown))
	// La victoire 1 donne 10, puis chaque victoire double le gain précédent.
	argent := 10 << (etage - 1)
	c.Money += argent
	messageSucces(fmt.Sprintf("+%d pièces d'or (bourse : %d).", argent, c.Money))
	switch etage {
	case 3: c.donnerButin("Plume divine", 2)
	case 4: c.donnerButin("Essence divine", 1)
	}
}

func (c *Character) donnerButin(nom string, quantite int) {
	libres := c.LimitInventaire - c.nombreObjets()
	if libres < quantite { quantite = libres }
	if quantite <= 0 { messageErreur("Inventaire plein : butin perdu : " + nom); return }
	c.addinventory(nom, quantite)
	messageSucces(fmt.Sprintf("Butin : %s x%d.", nom, quantite))
}
