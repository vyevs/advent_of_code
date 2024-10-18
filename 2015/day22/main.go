package main

import (
	"fmt"
	"slices"
	"time"
)

type boss struct {
	hp int

	poisonDur int
}

type player struct {
	hp       int
	mana     int
	manaUsed int

	shieldDur   int
	rechargeDur int
}

const (
	magicMissile = 53
	drain        = 73
	shield       = 113
	poison       = 173
	recharge     = 229

	shieldDur = 6
	shieldAmt = 7

	poisonDur = 6
	poisonDmg = 3

	rechargeDur = 5
	rechargeAmt = 101

	bossDmg = 10
)

type sim struct {
	spells       []int
	minSpellCost int
	hardMode     bool

	manaUsed []int
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %v\n", time.Since(start))
	}(time.Now())

	spells := []int{
		magicMissile,
		drain,
		shield,
		poison,
		recharge,
	}

	player := player{
		hp:   50,
		mana: 500,
	}
	boss := boss{
		hp: 71,
	}

	s := sim{
		spells:       spells,
		minSpellCost: magicMissile,
		hardMode:     false,
		manaUsed:     make([]int, 0, 1<<20),
	}

	s.run(player, boss)
}

func (s *sim) run(p player, b boss) {
	s.doTurn(p, b, true)
	fmt.Printf("least amount of mana used in a win: %d\n", slices.Min(s.manaUsed))
}

func (s *sim) doTurn(p player, b boss, pTurn bool) {
	if s.hardMode && pTurn {
		p.hp--

		if p.hp <= 0 {
			return
		}
	}

	p, b = applyEffects(p, b)

	if b.hp <= 0 {
		s.manaUsed = append(s.manaUsed, p.manaUsed)
		return
	}

	if pTurn {
		if p.mana < s.minSpellCost {
			return
		}

		for _, spell := range s.spells {
			if spell > p.mana {
				break
			}

			newP, newB, cast := castSpell(spell, p, b)
			if !cast {
				continue
			}

			if newB.hp <= 0 {
				s.manaUsed = append(s.manaUsed, newP.manaUsed)
				continue
			}

			s.doTurn(newP, newB, !pTurn)
		}

	} else {
		p = b.attack(p)

		if p.hp <= 0 {
			return
		}
		s.doTurn(p, b, !pTurn)
	}
}

func applyEffects(p player, b boss) (player, boss) {
	if b.poisonDur > 0 {
		b.hp -= poisonDmg

		b.poisonDur -= 1
	}

	if p.shieldDur > 0 {
		p.shieldDur -= 1
	}

	if p.rechargeDur > 0 {
		p.mana += rechargeAmt

		p.rechargeDur -= 1
	}

	return p, b
}

func (b boss) attack(p player) player {
	dmg := bossDmg

	if p.shieldDur > 0 {
		dmg = max(1, dmg-shieldAmt)
	}

	p.hp -= dmg

	return p
}

func castSpell(s int, p player, b boss) (player, boss, bool) {

	switch s {
	case magicMissile:
		{
			b.hp -= 4
		}
	case drain:
		{
			b.hp -= 2
			p.hp += 2
		}
	case shield:
		{
			if p.shieldDur > 0 {
				return p, b, false
			}
			p.shieldDur = shieldDur
		}
	case poison:
		{
			if b.poisonDur > 0 {
				return p, b, false
			}
			b.poisonDur = poisonDur
		}
	case recharge:
		{
			if p.rechargeDur > 0 {
				return p, b, false
			}
			p.rechargeDur = rechargeDur
		}
	}

	p.mana -= s
	p.manaUsed += s

	return p, b, true
}

func (p player) String() string {
	armor := 0
	if p.shieldDur > 0 {
		armor = shieldAmt
	}
	return fmt.Sprintf("Player has %d hit points, %d armor, %d mana", p.hp, armor, p.mana)
}

func (b boss) String() string {
	return fmt.Sprintf("Boss has %d hit points", b.hp)
}
