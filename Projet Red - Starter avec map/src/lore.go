package src

import "fmt"

func afficherLoreEntreeTour() {
	fmt.Println(`
À Athènes, on raconte que cette tour apparaît tous les mille ans.
Certains disent qu'elle fut construite pour divertir les Dieux.
D'autres affirment que ceux qui atteignent son terme peuvent devenir
leurs égaux.

Tu ne sais pas quelle histoire croire. Pourtant, tu es venu ici
avec une ambition que tu ne peux plus ignorer : devenir un Dieu.
`)
}

func afficherLoreChoixVoie(direction string) {
	switch direction {
	case "Monter":
		fmt.Println(`
Tu avances sur les marches baignées d'une lumière chaude.
L'air est calme, presque rassurant.

Tu repenses aux récits entendus à Athènes : au bout de cette voie,
un mortel pourrait gagner sa place parmi les Dieux de l'Olympe.
`)

	case "Descendre":
		fmt.Println(`
Tu t'engages dans l'escalier obscur. Quelques marches suffisent
pour que la lumière d'Athènes disparaisse derrière toi.

D'après les histoires racontées en ville, cette voie traverse
les profondeurs de la Terre et mène jusqu'aux Enfers.
`)
	}
}

func afficherLoreEtageTrois() {
	fmt.Println(`
Après le combat, tu remarques des noms gravés dans la pierre.
Ils appartiennent à ceux qui ont emprunté la Tour avant toi.

Tu ignores ce qu'ils sont devenus. Leurs noms, eux, sont restés.
`)
}

func afficherLoreDernierEtage(direction string) {
	switch direction {
	case "Monter":
		fmt.Println(`
La dernière porte s'ouvre sur un ciel éclatant.
Tu as entendu le nom de Zeus toute ta vie.
Cette fois, il se tient devant toi.
`)

	case "Descendre":
		fmt.Println(`
La dernière porte s'ouvre sur une vaste salle silencieuse.
Tu as entendu d'innombrables histoires sur Hadès.
Aucune ne t'avait préparé à le rencontrer.
`)
	}
}