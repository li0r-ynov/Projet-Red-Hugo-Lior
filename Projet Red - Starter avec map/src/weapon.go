package src

type Weapon struct {
	name []Winfos
}
type Winfos struct {
	name    string
	damage  int
	vitesse int
}

var offgear = map[string]Winfos{
	"CDP":               {name: "CDP", damage: 5, vitesse: 7},
	"Dagues":            {name: "Dagues", damage: 10, vitesse: 10},
	"Hache":             {name: "Hache", damage: 15, vitesse: 5},
	"Marteau de Guerre": {name: "Marteau de Guerre", damage: 17, vitesse: 0},
	"Épée":              {name: "Épée", damage: 13, vitesse: 7},
}
