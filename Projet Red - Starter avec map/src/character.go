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
	PotionPaqueAchetee      bool
	ArmeDivineObtenue       bool
	LimitInventaire         int
	UpgradeCount            int
	Damage                  int
	Vitesse                 int
	ArmeEquipee             string
	ArmuresEquipee          map[string]string
	DirectionTour           string
	EtageTour               int
	PalierPvAtteint         int
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
	c.EtageTour = 1
	c.PalierPvAtteint = 0
	c.DirectionTour = ""

	c.Inventaire = make(map[string]int)
	c.ArmuresEquipee = make(map[string]string)
}

func (c Character) displayInfo() {
	titreEcran("FICHE DU PERSONNAGE")

	sousTitre("IDENTITÉ")
	ligneInfo("Nom", c.Name)
	ligneInfo("Classe", c.Classe)
	ligneInfo("Titre", c.Level)
	ligneInfo("Renommée", fmt.Sprint(c.Renown))

	sousTitre("RESSOURCES")
	ligneInfo("Points de vie", fmt.Sprintf("%d/%d", c.Pv, c.PvMax))
	ligneInfo("Pièces d'or", fmt.Sprint(c.Money))
	ligneInfo("Inventaire", fmt.Sprintf("%d/%d", c.nombreObjets(), c.LimitInventaire))

	sousTitre("ÉQUIPEMENT")
	ligneInfo("Arme", equipmentName(c.ArmeEquipee))

	for _, emplacement := range []string{"Casque", "Plastron", "Bottes"} {
		ligneInfo(emplacement, equipmentName(c.ArmuresEquipee[emplacement]))
	}

	sousTitre("COMBAT")
	ligneInfo("Dégâts naturels", fmt.Sprint(c.Damage))
	ligneInfo("Vitesse naturelle", fmt.Sprint(c.Vitesse))

	degats, vitesse := c.cdpStats()
	ligneInfo("CDP", fmt.Sprintf("%d dégâts • %d vitesse", degats, vitesse))

	if c.ArmeEquipee != "" {
		degats, vitesse = c.weaponStats()
		ligneInfo("Attaque d'arme", fmt.Sprintf("%d dégâts • %d vitesse", degats, vitesse))
	}
}

func equipmentName(nom string) string {
	if nom == "" {
		return "Aucun"
	}
	return nom
}