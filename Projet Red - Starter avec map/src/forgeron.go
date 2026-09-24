package src

import "fmt"

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

func (c *Character) MenuForgeron() {
	fmt.Println("\nLe forgeron vous présente ses plans de fabrication.")

	for {
		fmt.Println("\n=== Forgeron ===")

		for i, plan := range plansFabrication {
			fmt.Printf("%d - %s : ", i+1, plan.Armure)

			for _, materiau := range []string{
				"Fer", "Cuir", "Plume divine", "Essence divine",
			} {
				if quantite := plan.Materiaux[materiau]; quantite > 0 {
					fmt.Printf("%s x%d  ", materiau, quantite)
				}
			}

			fmt.Println()
		}

		fmt.Println("0 - Retourner au marché")
		fmt.Print("Équipement à fabriquer : ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println("Saisie invalide.")
			return
		}

		if choix == 0 {
			return
		}

		if choix < 1 || choix > len(plansFabrication) {
			fmt.Println("Choix invalide.")
			continue
		}

		c.fabriquerArmure(plansFabrication[choix-1])
	}
}

func (c *Character) fabriquerArmure(plan PlanFabrication) {
	for materiau, quantite := range plan.Materiaux {
		if c.Inventaire[materiau] < quantite {
			fmt.Printf(
				"Matériau manquant : %s x%d (vous en avez %d).\n",
				materiau, quantite, c.Inventaire[materiau],
			)
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
		fmt.Println("Votre inventaire sera plein après la fabrication.")
		return
	}

	for materiau, quantite := range plan.Materiaux {
		c.Inventaire[materiau] -= quantite

		if c.Inventaire[materiau] == 0 {
			delete(c.Inventaire, materiau)
		}
	}

	c.addinventory(plan.Armure, 1)
	fmt.Printf("%s a été fabriqué !\n", plan.Armure)
	c.equipArmor(plan.Armure)
}