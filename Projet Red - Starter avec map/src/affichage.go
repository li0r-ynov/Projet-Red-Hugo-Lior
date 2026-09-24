package src

import "fmt"

const (
	dore    = "\033[38;2;212;175;55m"
	crimson = "\033[38;2;220;20;60m"
	bleu    = "\033[38;2;65;105;180m"
	blanc   = "\033[97m"
	gris    = "\033[90m"
	reset   = "\033[0m"
)

func titreEcran(titre string) {
	fmt.Printf("\n%s╔══════════════════════════════════════════════╗%s\n", dore, reset)
	fmt.Printf("%s║ %-44s ║%s\n", dore, titre, reset)
	fmt.Printf("%s╚══════════════════════════════════════════════╝%s\n", dore, reset)
}

func sousTitre(titre string) {
	fmt.Printf("\n%s◆ %s%s\n", crimson, titre, reset)
}

func option(touche, description string) {
	fmt.Printf("  %s[%s]%s %s%s%s\n", dore, touche, reset, bleu, description, reset)
}

func separateur() {
	fmt.Printf("%s────────────────────────────────────────────────%s\n", gris, reset)
}

func messageSucces(texte string) {
	fmt.Printf("%s✦ %s%s\n", dore, texte, reset)
}

func messageErreur(texte string) {
	fmt.Printf("%s✦ %s%s\n", crimson, texte, reset)
}