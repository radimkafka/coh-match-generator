package main

import (
	. "coh-match-generator/internal"
	"fmt"
	"math/rand"
)

func main() {
	match := generateMatch()
	printMatch(&match)
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
		Player:   axis,
		Faction:  axisFaction.Name,
		Doctrine: axisFaction.Doctrines[rand.Intn(3)],
	}
	alliesOutput := MatchItem{
		Player:   allies,
		Faction:  alliesFaction.Name,
		Doctrine: alliesFaction.Doctrines[rand.Intn(3)],
	}

	matchInfo := MatchInfo{
		MapName: Maps[rand.Intn(len(Maps)-1)],
		Mode:    Modes[rand.Intn(2)],
		Axis:    axisOutput,
		Allies:  alliesOutput,
	}
	return matchInfo
}

func printMatch(match *MatchInfo) {
	fmt.Println(match.MapName, "-", match.Mode)
	fmt.Println(match.Allies.Player, "-", match.Allies.Faction, "-", match.Allies.Doctrine)
	fmt.Println(match.Axis.Player, "-", match.Axis.Faction, "-", match.Axis.Doctrine)
}
