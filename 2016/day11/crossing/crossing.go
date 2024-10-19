package main

import (
	"fmt"
	"slices"
	"time"
)

const (
	farmer   = 0
	wolf     = 1
	goat     = 2
	cabbage  = 3
	allItems = 4
)

var initialState = [4]byte{0, 0, 0, 0}
var desiredState = [4]byte{1, 1, 1, 1}

func wolf_cant_eat_goat(state [4]byte) bool {
	return state[farmer] == state[wolf] || state[wolf] != state[goat]
}

func goat_cant_eat_cabbage(state [4]byte) bool {
	return state[farmer] == state[goat] || state[goat] != state[cabbage]
}

func valid_state(state [4]byte) bool {
	return wolf_cant_eat_goat(state) && goat_cant_eat_cabbage(state)
}

func get_next_states(state [4]byte) [][4]byte {
	nextStates := make([][4]byte, 0, allItems)
	for other := range allItems {
		// Farmer and item to take need to be on the same side.
		if state[other] == state[farmer] {
			newState := state

			newState[farmer] ^= 1
			if other != farmer {
				newState[other] ^= 1
			}

			if valid_state(newState) {
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

type crosser struct {
	states [][4]byte
}

func (c *crosser) cross(state [4]byte) bool {
	c.states = append(c.states, state)

	if state == desiredState {
		return true
	}

	nextStates := get_next_states(state)
	for _, nextState := range nextStates {
		if !slices.Contains(c.states, nextState) {
			if c.cross(nextState) {
				return true
			}
			c.states = c.states[:len(c.states)-1]
		}
	}

	return false
}

func cross_bfs(start [4]byte) [][4]byte {
	type node struct {
		state  [4]byte
		parent *node
	}

	collect_result := func(end *node) [][4]byte {
		path := make([][4]byte, 0, 16)
		for cur := end; cur != nil; cur = cur.parent {
			path = append(path, cur.state)
		}
		slices.Reverse(path)
		return path
	}

	queue := make([]node, 0, 16)
	queue = append(queue, node{state: start})

	seen := make(map[[4]byte]bool, 16)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		seen[n.state] = true

		if n.state == desiredState {
			return collect_result(&n)
		}

		next_states := get_next_states(n.state)
		for _, ns := range next_states {
			if !seen[ns] {
				nsNode := node{
					state:  ns,
					parent: &n,
				}
				queue = append(queue, nsNode)
			}
		}
	}

	return nil
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %s\n", time.Since(start))
	}(time.Now())

	path := cross_bfs(initialState)
	if path == nil {
		fmt.Println("didn't make it")
	} else {
		fmt.Printf("made it in %d steps\n", len(path)-1)
		for _, s := range path {
			fmt.Println(s)
		}
	}
}
