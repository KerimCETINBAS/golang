package main

// custom type based on string
type Race string

// enums associated with Race
const (
	Hyur    Race = "Hyur"
	Lalafel Race = "Lalafel"
	Miqote  Race = "Miqo'te"
)

func parseRace(value string) Race {

	r := map[string]Race{
		"Hyur":    Race(Hyur),
		"Lalafel": Race(Lalafel),
		"Miqote":  Race(Miqote),
	}

	for v, race := range r {

		if v == value {
			return race

		}
	}

	return ""
}

type hero struct {
	name string
	race Race
}

func createHero(name string, race Race) hero {

	return hero{
		name: name,
		race: Race(race),
	}
}
