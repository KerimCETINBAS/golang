package main

// custom type based on string
type Race string

// enums associated with Race
const (
	Hyur    Race = "Hyur"
	Lalafel Race = "Lalafel"
	Miqote  Race = "Miqo'te"
)

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

func (h *hero) rename(name string) {
	h.name = name
}

func (h hero) getRace() Race {

	return h.race
}
