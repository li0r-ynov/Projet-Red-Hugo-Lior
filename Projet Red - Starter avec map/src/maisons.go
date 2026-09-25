package src

import "fmt"

func (c *Character) MenuMaisonsDieux() {
	for {
		titreEcran("MAISONS DES DIEUX")
		option("1", "Maison de Zeus")
		option("2", "Maison d'Hadès")
		option("3", "Maison de Poséidon")
		option("0", "Retour à Athènes")
		fmt.Print("Votre choix : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch choix {
		case 0:
			return
		case 1:
			c.MaisonZeus()
		case 2:
			c.MaisonHades()
		case 3:
			c.MaisonPoseidon()
		default:
			messageErreur("Choix invalide.")
		}
	}
}

func (c *Character) offrirArmeDivine(dieu, arme string) {
	titreEcran("MAISON DE " + dieu)
	ligneInfo("Arme légendaire", arme)

	if c.PalierPvAtteint < 3 {
		messageErreur("Atteignez d'abord le titre Héros ou Démon.")
		return
	}
	if c.ArmeDivineObtenue {
		messageErreur("Vous avez déjà obtenu une arme divine.")
		return
	}
	if c.Inventaire[PotionPaque] <= 0 {
		messageErreur("Vous n'avez pas la potion de Pâques.")
		return
	}

	option("1", "Offrir la potion et recevoir l'arme")
	option("0", "Retour")
	fmt.Print("Votre choix : ")

	var choix int
	if _, err := fmt.Scan(&choix); err != nil {
		messageErreur("Saisie invalide.")
		return
	}
	if choix != 1 {
		return
	}

	// Consommer la potion libère une place pour l'arme.
	c.removeInventory(PotionPaque, 1)
	c.addinventory(arme, 1)
	c.ArmeDivineObtenue = true
	messageSucces("Vous recevez : " + arme)
}