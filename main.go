package main

import (
	"fmt"
	"math/rand"
)

type Fraction struct {
	name      string
	doctrines []string
}

type MatchItem = struct {
	player   string
	faction  string
	doctrine string
}

type MatchInfo = struct {
	allies  MatchItem
	axis    MatchItem
	mapName string
	mode    string
}

var Maps = []string{
	"Angoville",
	"Beach Assault",
	"Beaux Lowlands",
	"Bernieres-Sur-Mer",
	"Best",
	"Carpiquet",
	"Duclair",
	"Egletons",
	"Flooded Plains",
	"Hinderdam",
	"Industrial Riverbed",
	"Langres",
	// "Lyon",
	"Ruins of Rouen",
	"Semois",
	"St. Mere Dumont",
	"Sturzdorf",
	"Verrieres Ridge",
	"Verrieres Ridge - No Bunkers",
	"Wrecked Train",
	"Wolfheze (2v2)",
}

var Modes = []string{"Victory Point Control", "Annihilation"}

var Americans = Fraction{
	name:      "Americans",
	doctrines: []string{"Infantry Company", "Airborne Company", "Armor Company"},
}
var British = Fraction{
	name: "British",
	doctrines: []string{
		"Royal Engineers Support",
		"Royal Commandos Support",
		"Royal Artillery Support",
	},
}

var PanzerElite = Fraction{
	name: "Panzer Elite",
	doctrines: []string{
		"Scorched Earth Tactics",
		"Luftwaffe Tactics",
		"Tank Destroyer Tactics",
	},
}

var Wehrmacht = Fraction{
	name:      "Wehrmacht",
	doctrines: []string{"Defensive Doctrine", "Blitzkrieg Doctrine", "Terror Doctrine"},
}

var AlliesFactions = []Fraction{Americans}
var AxisFactions = []Fraction{Wehrmacht, PanzerElite}

func main() {
	match := generateMatch()
	fmt.Println(match.mapName, "-", match.mode)
	fmt.Println(match.allies.player, "-", match.allies.faction, "-", match.allies.doctrine)
	fmt.Println(match.axis.player, "-", match.axis.faction, "-", match.axis.doctrine)
}

func generateMatch() MatchInfo {
	p1 := "Já"
	p2 := "Pája"

	allies := ""
	axis := ""

	if rand.Intn(2) == 0 {
		allies = p1
		axis = p2
	} else {
		allies = p2
		axis = p1
	}
	alliesFaction := AlliesFactions[rand.Intn(len(AlliesFactions))]
	axisFaction := AxisFactions[rand.Intn(len(AxisFactions))]

	axisOutput := MatchItem{
		player:   axis,
		faction:  axisFaction.name,
		doctrine: axisFaction.doctrines[rand.Intn(3)],
	}
	alliesOutput := MatchItem{
		player:   allies,
		faction:  alliesFaction.name,
		doctrine: alliesFaction.doctrines[rand.Intn(3)],
	}

	matchInfo := MatchInfo{
		mapName: Maps[rand.Intn(len(Maps)-1)],
		mode:    Modes[rand.Intn(2)],
		axis:    axisOutput,
		allies:  alliesOutput,
	}
	return matchInfo
}
