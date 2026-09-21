package main

import "fmt"

func (c *Character) TowerTravelDisplay(text, textOption1, textOption2 string) int {
	fmt.Println(text)
	fmt.Println("\t 1 - ", textOption1)
	fmt.Println("\t 2 - ", textOption2)

	fmt.Print("Votre choix : ")
	var LoreChoice int
	fmt.Scan(&LoreChoice)
}
