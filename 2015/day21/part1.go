package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	cost  int
	dmg   int
	armor int
}

var (
	dagger     = item{cost: 8, dmg: 4}
	shortsword = item{cost: 10, dmg: 5}
	warhammer  = item{cost: 25, dmg: 6}
	longsword  = item{cost: 40, dmg: 7}
	greataxe   = item{cost: 74, dmg: 8}
	weapons    = []item{dagger, shortsword, warhammer, longsword, greataxe}

	leather    = item{cost: 13, armor: 1}
	chainmail  = item{cost: 31, armor: 2}
	splintmail = item{cost: 53, armor: 3}
	bandedmail = item{cost: 75, armor: 4}
	platemail  = item{cost: 102, armor: 5}
	armor      = []item{item{}, leather, chainmail, splintmail, bandedmail, platemail}

	ringDmg1 = item{cost: 25, dmg: 1}
	ringDmg2 = item{cost: 50, dmg: 2}
	ringDmg3 = item{cost: 100, dmg: 3}
	ringDef1 = item{cost: 20, armor: 1}
	ringDef2 = item{cost: 40, armor: 2}
	ringDef3 = item{cost: 80, armor: 3}
	rings    = []item{item{}, item{}, ringDmg1, ringDmg2, ringDmg3, ringDef1, ringDef2, ringDef3}
)

type player struct {
	name  string
	hp    int
	dmg   int
	armor int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	boss := parseBoss(f)
	f.Close()

	player := player{name: "player", hp: 100}

	bestCost := -1
	var bestSetup []item
	for _, weapon := range weapons {
		for _, armor := range armor {
			for firstRingIndex, ring1 := range rings {
				for secondRingIndex := firstRingIndex + 1; secondRingIndex < len(rings); secondRingIndex++ {

					ring2 := rings[secondRingIndex]

					items := []item{weapon, armor, ring1, ring2}

					player := applyItems(player, items)

					winner := fight(player, boss)

					if winner.name == "player" {
						cost := totalCost(items)
						if bestCost == -1 || cost < bestCost {
							bestCost = cost
							bestSetup = items
						}
					}
				}
			}
		}
	}

	fmt.Printf("can win only spending %d gold\n", bestCost)
	fmt.Printf("setup: %+v\n", bestSetup)

	var worstCost int
	var worstSetup []item
	for _, weapon := range weapons {
		for _, armor := range armor {
			for firstRingIndex, ring1 := range rings {
				for secondRingIndex := firstRingIndex + 1; secondRingIndex < len(rings); secondRingIndex++ {

					ring2 := rings[secondRingIndex]

					items := []item{weapon, armor, ring1, ring2}

					player := applyItems(player, items)

					winner := fight(player, boss)

					if winner.name == "boss" {
						cost := totalCost(items)
						if cost > worstCost {
							worstCost = cost
							worstSetup = items
						}
					}
				}
			}
		}
	}

	fmt.Printf("can lose spending %d gold\n", worstCost)
	fmt.Printf("setup: %+v\n", worstSetup)
}

func applyItems(p player, items []item) player {
	for _, item := range items {
		p.dmg += item.dmg
		p.armor += item.armor
	}
	return p
}

func totalCost(items []item) int {
	var cost int
	for _, item := range items {
		cost += item.cost
	}
	return cost
}

func fight(p1, p2 player) player {
	p1Turn := true
	for {
		if p1Turn {
			dmg := p1.dmg - p2.armor
			if dmg < 1 {
				dmg = 1
			}

			p2.hp -= dmg

			if p2.hp <= 0 {
				return p1
			}

		} else {
			dmg := p2.dmg - p1.armor
			if dmg < 1 {
				dmg = 1
			}

			p1.hp -= dmg

			if p1.hp <= 0 {
				return p2
			}
		}

		p1Turn = !p1Turn
	}
}

func parseBoss(r io.Reader) player {
	scanner := bufio.NewScanner(r)

	boss := player{
		name: "boss",
	}

	{
		scanner.Scan()
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		hpStr := tokens[2]
		hp, err := strconv.Atoi(hpStr)
		if err != nil {
			log.Fatalf("error parsing hp: %v", err)
		}
		boss.hp = hp
	}

	{
		scanner.Scan()
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		dmgStr := tokens[1]
		dmg, err := strconv.Atoi(dmgStr)
		if err != nil {
			log.Fatalf("error parsing hp: %v", err)
		}
		boss.dmg = dmg
	}

	{
		scanner.Scan()
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		armorStr := tokens[1]
		armor, err := strconv.Atoi(armorStr)
		if err != nil {
			log.Fatalf("error parsing hp: %v", err)
		}
		boss.armor = armor
	}

	return boss
}
