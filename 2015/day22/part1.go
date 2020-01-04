package main

type boss struct {
	hp  int
	dmg int

	poisonDuration int
}

type player struct {
	hp   int
	mana int

	shieldDuration   int
	rechargeDuration int
}

func main() {
	player := player{hp: 50, mana: 500}
	boss := boss{hp: 71, dmg: 10}

}
