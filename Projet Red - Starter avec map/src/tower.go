package src

import "fmt"

func (c *Character) TowerTravelDisplay(text, textOption1, textOption2, textOption3, textOption4, textOption5 string) int {

	fmt.Println(text)
	fmt.Println("\t 1 - ", textOption1)
	fmt.Println("\t 2 - ", textOption2)
	fmt.Println("\t 3 - ", textOption3)
	fmt.Println("\t 4 - ", textOption4)
	fmt.Println("\t 0 - ", textOption5)
	fmt.Print("Votre choix : ")
	var LoreChoice int
	fmt.Scan(&LoreChoice)
	return LoreChoice
}
