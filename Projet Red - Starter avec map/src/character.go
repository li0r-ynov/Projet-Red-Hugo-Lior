package src

import "fmt"

type Character struct {
	Name                    string
	Classe                  string
	PvMax                   int
	Pv                      int
	Inventaire              map[string]int
	Money                   int
	Level                   string
	Renown                  int
	PotionGratuiteRecuperee bool
	LimitInventaire         int
	Damage                  int
	Vitesse                 int
	ArmeEquipee             string
	ArmuresEquipee          map[string]string

	DirectionTour  string
	EtageTour      int
	PalierPvAtteint int
}

func (c *Character) initCharacter() {
	c.LimitInventaire = 10

	switch c.Classe {
	case "Athénien":
		c.PvMax = 50
		c.Damage = 15
	case "Sparte":
		c.PvMax = 60
		c.Damage = 10
	}

	c.Pv = c.PvMax / 2
	c.Money = 10
	c.Vitesse = 10
	c.Level = "Civil"
	c.Renown = 0
	c.PalierPvAtteint = 0
	c.DirectionTour = ""
	c.EtageTour = 1

	c.Inventaire = make(map[string]int)
	c.ArmuresEquipee = make(map[string]string)
	c.ArmeEquipee = ""
}

func (c Character) displayInfo() {
	totalObjets := 0
	for _, quantite := range c.Inventaire {
		totalObjets += quantite
	}

	titreEcran("FICHE DU PERSONNAGE")

	sousTitre("IDENTITÉ")
	ligneInfo("Nom", c.Name)
	ligneInfo("Classe", c.Classe)
	ligneInfo("Titre", c.Level)
	ligneInfo("Renommée", fmt.Sprintf("%d", c.Renown))

	sousTitre("RESSOURCES")
	ligneInfo("Points de vie", fmt.Sprintf("%d/%d", c.Pv, c.PvMax))
	ligneInfo("Pièces d'or", fmt.Sprintf("%d", c.Money))
	ligneInfo("Inventaire", fmt.Sprintf("%d/%d", totalObjets, c.LimitInventaire))

	sousTitre("ÉQUIPEMENT")
	ligneInfo("Arme", equipmentName(c.ArmeEquipee))
	ligneInfo("Casque", equipmentName(c.ArmuresEquipee["Casque"]))
	ligneInfo("Plastron", equipmentName(c.ArmuresEquipee["Plastron"]))
	ligneInfo("Bottes", equipmentName(c.ArmuresEquipee["Bottes"]))

	sousTitre("COMBAT")
	ligneInfo("Dégâts naturels", fmt.Sprintf("%d", c.Damage))
	ligneInfo("Vitesse naturelle", fmt.Sprintf("%d", c.Vitesse))

	degatsCDP, vitesseCDP := c.cdpStats()
	ligneInfo("CDP", fmt.Sprintf("%d dégâts • %d vitesse", degatsCDP, vitesseCDP))

	if c.ArmeEquipee != "" {
		degatsArme, vitesseArme := c.weaponStats()
		ligneInfo(
			"Attaque d'arme",
			fmt.Sprintf("%d dégâts • %d vitesse", degatsArme, vitesseArme),
		)
	}

	separateur()
}

func ligneInfo(nom, valeur string) {
	fmt.Printf("  %s%-19s%s %s%s%s\n",
		bleu, nom, reset,
		blanc, valeur, reset)
}

func equipmentName(nom string) string {
	if nom == "" {
		return "Aucun"
	}

	return nom
}