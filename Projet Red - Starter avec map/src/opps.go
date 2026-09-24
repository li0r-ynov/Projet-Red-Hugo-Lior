package src

type opps struct {
	Name   string
	PvMax  int
	Pv     int
	Level  string
	Damage int
}

var oppsdef = map[string]opps{
	"Combattant": {Name: "Combattant", PvMax: 60, Pv: 60, Damage: 15},
	"Guerrier":   {Name: "Guerrier", PvMax: 70, Pv: 70, Damage: 18},
	"Civil":     {Name: "Civile", PvMax: 50, Pv: 50, Damage: 10},
	"Bandit":     {Name: "Bandit", PvMax: 60, Pv: 60, Damage: 15},
	"Mercenaire": {Name: "Mercenaire", PvMax: 70, Pv: 70, Damage: 18},
	"Héro":       {Name: "Héro", PvMax: 100, Pv: 100, Damage: 22},
	"Démon":      {Name: "Démon", PvMax: 100, Pv: 100, Damage: 22},
	"Demi-Dieu":  {Name: "Demi-Dieu", PvMax: 120, Pv: 120, Damage: 25},
	"Zeus":       {Name: "Zeus", PvMax: 150, Pv: 150, Damage: 30},
	"Hadès":      {Name: "Hadès", PvMax: 150, Pv: 150, Damage: 30},
}
