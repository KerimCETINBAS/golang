package main

type hero struct {
	name  string
	items map[int]string
	race  string
}

func newHero(name string, race string) hero {
	return hero{
		name: name,
		race: race,
	}
}
