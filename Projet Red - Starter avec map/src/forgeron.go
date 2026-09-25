package src

import (
	"fmt"
	"strings"
)

type PlanFabrication struct {
	Armure    string
	Materiaux map[string]int
}

var plansFabrication = []PlanFabrication{
	{Armure: "Casque de Gladiateur", Materiaux: map[string]int{"Fer": 1, "Cuir": 1}},
	{Armure: "Plastron de Gladiateur", Materiaux: map[string]int{"Fer": 2, "Cuir": 1}},
	{Armure: "Bottes de Gladiateur", Materiaux: map[string]int{"Fer": 1, "Cuir": 1}},
	{Armure: "Bottes de Hermès", Materiaux: map[string]int{"Cuir": 1, "Plume divine": 1}},
	{Armure: "Ailes d'Icare", Materiaux: map[string]int{"Cuir": 1, "Plume divine": 2}},
	{Armure: "Casque d'Arès", Materiaux: map[string]int{"Fer": 1, "Essence divine": 1}},
}

func afficherPlansFabrication() {
	titreEcran("FORGERON")
	sousTitre("PLANS DE FABRICATION")

	for i, plan := range plansFabrication {
		fmt.Printf(
			"  %s[%d]%s %s%s%s\n      ",
			dore, i+1, reset, bleu, plan.Armure, reset,
		)

		premier := true
		for _, nom := range []string{
			"Fer", "Cuir", "Plume divine", "Essence divine",
		} {
			if quantite := plan.Materiaux[nom]; quantite > 0 {
				if !premier {
					fmt.Print(" • ")
				}
				fmt.Printf("%s%s x%d%s", blanc, nom, quantite, reset)
				premier = false
			}
		}
		fmt.Println()
	}

	option("M", "Revoir les plans")
	option("0", "Retour au marché")
}

func (c *Character) MenuForgeron() {
	afficherPlansFabrication()

	for {
		fmt.Print("\nArmure à fabriquer (M : plans, 0 : retour) : ")

		var choix string
		if _, err := fmt.Scan(&choix); err != nil {
			messageErreur("Saisie invalide.")
			return
		}

		switch strings.ToUpper(choix) {
		case "0":
			return
		case "M":
			afficherPlansFabrication()
		case "1", "2", "3", "4", "5", "6":
			numero := int(choix[0] - '1')
			c.fabriquerArmure(plansFabrication[numero])
		default:
			messageErreur("Choix invalide.")
		}
	}
}

func (c *Character) fabriquerArmure(plan PlanFabrication) {
	for nom, quantite := range plan.Materiaux {
		if c.Inventaire[nom] < quantite {
			messageErreur(fmt.Sprintf(
				"Il manque %s : %d/%d.",
				nom, c.Inventaire[nom], quantite,
			))
			return
		}
	}

	consommes := 0
	for _, quantite := range plan.Materiaux {
		consommes += quantite
	}

	if c.nombreObjets()-consommes+1 > c.LimitInventaire {
		messageErreur("Inventaire plein après fabrication.")
		return
	}

	for nom, quantite := range plan.Materiaux {
		c.removeInventory(nom, quantite)
	}

	c.addinventory(plan.Armure, 1)

	switch plan.Armure {
	case "Bottes de Gladiateur", "Bottes de Hermès", "Ailes d'Icare":
		messageSucces(plan.Armure + " ont été fabriquées !")
	default:
		messageSucces(plan.Armure + " a été fabriqué !")
	}

	c.equipArmor(plan.Armure)
}