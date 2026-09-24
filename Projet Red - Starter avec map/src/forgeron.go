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
	{
		Armure: "Casque de Gladiateur",
		Materiaux: map[string]int{
			"Fer":  1,
			"Cuir": 1,
		},
	},
	{
		Armure: "Plastron de Gladiateur",
		Materiaux: map[string]int{
			"Fer":  2,
			"Cuir": 1,
		},
	},
	{
		Armure: "Bottes de Gladiateur",
		Materiaux: map[string]int{
			"Fer":  1,
			"Cuir": 1,
		},
	},
	{
		Armure: "Bottes de Hermès",
		Materiaux: map[string]int{
			"Cuir":         1,
			"Plume divine": 1,
		},
	},
	{
		Armure: "Ailes d'Icare",
		Materiaux: map[string]int{
			"Cuir":         1,
			"Plume divine": 2,
		},
	},
	{
		Armure: "Casque d'Arès",
		Materiaux: map[string]int{
			"Fer":            1,
			"Essence divine": 1,
		},
	},
}

func afficherPlansFabrication() {
	titreEcran("FORGERON")
	fmt.Printf("  %sLe forgeron vous présente ses plans.%s\n", blanc, reset)
	separateur()
	sousTitre("PLANS DE FABRICATION")

	for i, plan := range plansFabrication {
		fmt.Printf("  %s[%d]%s %s%s%s\n",
			dore, i+1, reset,
			bleu, plan.Armure, reset)

		fmt.Print("      ")
		premier := true

		for _, nom := range []string{
			"Fer", "Cuir", "Plume divine", "Essence divine",
		} {
			quantite := plan.Materiaux[nom]
			if quantite == 0 {
				continue
			}

			if !premier {
				fmt.Print("  •  ")
			}

			fmt.Printf("%s%s x%d%s", blanc, nom, quantite, reset)
			premier = false
		}

		fmt.Println()
	}

	fmt.Println()
	option("M", "Revoir les plans")
	option("0", "Retourner au marché")
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
			messageErreur("Choix invalide. Entrez un numéro, M ou 0.")
		}
	}
}

func (c *Character) fabriquerArmure(plan PlanFabrication) {
	for nom, necessaire := range plan.Materiaux {
		possede := c.Inventaire[nom]

		if possede < necessaire {
			messageErreur(fmt.Sprintf(
				"Il manque %s : %d/%d.",
				nom, possede, necessaire,
			))
			return
		}
	}

	totalObjets := 0
	for _, quantite := range c.Inventaire {
		totalObjets += quantite
	}

	materiauxConsommes := 0
	for _, quantite := range plan.Materiaux {
		materiauxConsommes += quantite
	}

	if totalObjets-materiauxConsommes+1 > c.LimitInventaire {
		messageErreur("Votre inventaire sera plein après la fabrication.")
		return
	}

	for nom, quantite := range plan.Materiaux {
		c.Inventaire[nom] -= quantite
		if c.Inventaire[nom] == 0 {
			delete(c.Inventaire, nom)
		}
	}

	c.addinventory(plan.Armure, 1)

	switch plan.Armure {
	case "Bottes de Gladiateur", "Bottes de Hermès":
		messageSucces(plan.Armure + " ont été fabriquées !")
	case "Ailes d'Icare":
		messageSucces(plan.Armure + " ont été fabriquées !")
	default:
		messageSucces(plan.Armure + " a été fabriqué !")
	}

	c.equipArmor(plan.Armure)
}