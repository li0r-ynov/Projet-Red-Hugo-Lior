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
	"CDP":                            {name: "CDP", damage: 5, vitesse: 6},
	"Dagues de l'Assassin":           {name: "Dagues de l'Assassin", damage: 10, vitesse: 10},
	"Hache de Viking":                {name: "Hache de Viking", damage: 15, vitesse: 5},
	"Marteau de Guerre du Martelier": {name: "Marteau de Guerre du Martelier", damage: 17, vitesse: 0},
	"Glaive du Légionnaire":          {name: "Glaive du Légionnaire", damage: 13, vitesse: 7},
}
