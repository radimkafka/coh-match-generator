package data

type Fraction struct {
	Name      string
	Doctrines []string
}

type MatchItem = struct {
	Player   string
	Faction  string
	Doctrine string
}

type MatchInfo = struct {
	Allies  MatchItem
	Axis    MatchItem
	MapName string
	Mode    string
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
	Name:      "Americans",
	Doctrines: []string{"Infantry Company", "Airborne Company", "Armor Company"},
}
var British = Fraction{
	Name: "British",
	Doctrines: []string{
		"Royal Engineers Support",
		"Royal Commandos Support",
		"Royal Artillery Support",
	},
}

var PanzerElite = Fraction{
	Name: "Panzer Elite",
	Doctrines: []string{
		"Scorched Earth Tactics",
		"Luftwaffe Tactics",
		"Tank Destroyer Tactics",
	},
}

var Wehrmacht = Fraction{
	Name:      "Wehrmacht",
	Doctrines: []string{"Defensive Doctrine", "Blitzkrieg Doctrine", "Terror Doctrine"},
}

var AlliesFactions = []Fraction{Americans, British}
var AxisFactions = []Fraction{Wehrmacht, PanzerElite}
