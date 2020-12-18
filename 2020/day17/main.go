package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	initialState := readInitialState()

	fmt.Println("Part 1:")
	doPart1(initialState)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(initialState)
}

func doPart1(activeCells [][2]int) {
	initialState := make(map[coord3]struct{}, len(activeCells))
	for _, c := range activeCells {
		c3 := coord3{
			x: c[0],
			y: c[1],
		}
		initialState[c3] = struct{}{}
	}

	state := initialState
	for i := 0; i < 6; i++ {
		state = nextState3(state)
	}

	var nActive int
	for range state {
		nActive++
	}

	fmt.Printf("After 6 cycles, %d cubes are active\n", nActive)
}

func nextState3(state map[coord3]struct{}) map[coord3]struct{} {
	next := make(map[coord3]struct{}, len(state)*2)

	minP := minPoint3(state)
	maxP := maxPoint3(state)

	for x := minP.x - 1; x <= maxP.x+1; x++ {
		for y := minP.y - 1; y <= maxP.y+1; y++ {
			for z := minP.z - 1; z <= maxP.z+1; z++ {
				c := coord3{
					x: x,
					y: y,
					z: z,
				}
				_, isActive := state[c]
				nNeighbors := nActiveNeighbors3(c, state)

				if isActive {
					if nNeighbors == 2 || nNeighbors == 3 {
						next[c] = struct{}{}
					}
				} else {
					if nNeighbors == 3 {
						next[c] = struct{}{}
					}
				}
			}
		}
	}

	return next
}

func nActiveNeighbors3(c coord3, state map[coord3]struct{}) int {
	var n int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				// do not want to consider the coord itself so at least one of dx, dy, dz need to != 0
				if dx != 0 || dy != 0 || dz != 0 {
					neighborC := coord3{
						c.x + dx,
						c.y + dy,
						c.z + dz,
					}
					_, isActive := state[neighborC]
					if isActive {
						n++
					}
				}
			}
		}
	}
	return n
}

func nActiveNeighbors4(c coord4, state map[coord4]struct{}) int {
	var n int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					// do not want to consider the coord itself so at least one of dx, dy, dz, dw need to != 0
					if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
						neighborC := coord4{
							c.x + dx,
							c.y + dy,
							c.z + dz,
							c.w + dw,
						}
						_, isActive := state[neighborC]
						if isActive {
							n++
						}
					}
				}
			}
		}
	}
	return n
}

func nextState4(state map[coord4]struct{}) map[coord4]struct{} {
	next := make(map[coord4]struct{}, len(state)*2)

	minP := minPoint4(state)
	maxP := maxPoint4(state)

	for x := minP.x - 1; x <= maxP.x+1; x++ {
		for y := minP.y - 1; y <= maxP.y+1; y++ {
			for z := minP.z - 1; z <= maxP.z+1; z++ {
				for w := minP.w - 1; w <= maxP.w+1; w++ {
					c := coord4{
						x: x,
						y: y,
						z: z,
						w: w,
					}
					_, isActive := state[c]
					nNeighbors := nActiveNeighbors4(c, state)

					if isActive {
						if nNeighbors == 2 || nNeighbors == 3 {
							next[c] = struct{}{}
						}
					} else {
						if nNeighbors == 3 {
							next[c] = struct{}{}
						}
					}
				}

			}
		}
	}

	return next
}

func doPart2(activeCells [][2]int) {
	initialState := make(map[coord4]struct{}, len(activeCells))
	for _, c := range activeCells {
		c3 := coord4{
			x: c[0],
			y: c[1],
		}
		initialState[c3] = struct{}{}
	}

	state := initialState
	for i := 0; i < 6; i++ {
		state = nextState4(state)
	}

	var nActive int
	for range state {
		nActive++
	}

	fmt.Printf("After 6 cycles, %d cubes are active\n", nActive)
}

type coord4 struct {
	x, y, z, w int
}

type coord3 struct {
	x, y, z int
}

// returns a slice of all the active (x, y) coordinates in a plane where z == 0 && w == 0
func readInitialState() [][2]int {
	scanner := bufio.NewScanner(os.Stdin)

	initialState := make([][2]int, 0, 4096)

	var y int
	for scanner.Scan() {

		line := scanner.Text()

		var x int
		for _, r := range line {
			if r == '#' {
				coord := [2]int{x, y}
				initialState = append(initialState, coord)
			}
			x++
		}

		y++
	}

	return initialState
}

func minPoint3(state map[coord3]struct{}) coord3 {
	maxInt := 100000000000000

	minX := maxInt
	for c := range state {
		if c.x < minX {
			minX = c.x
		}
	}

	minY := maxInt
	for c := range state {
		if c.y < minY {
			minY = c.y
		}
	}

	minZ := maxInt
	for c := range state {
		if c.z < minZ {
			minZ = c.z
		}
	}

	return coord3{
		x: minX,
		y: minY,
		z: minZ,
	}
}

func maxPoint3(state map[coord3]struct{}) coord3 {
	minInt := -100000000000000

	maxX := minInt
	for c := range state {
		if c.x > maxX {
			maxX = c.x
		}
	}

	maxY := minInt
	for c := range state {
		if c.y > maxY {
			maxY = c.y
		}
	}

	maxZ := minInt
	for c := range state {
		if c.z > maxZ {
			maxZ = c.z
		}
	}

	return coord3{
		x: maxX,
		y: maxY,
		z: maxZ,
	}
}

func minPoint4(state map[coord4]struct{}) coord4 {
	maxInt := 100000000000000

	minX := maxInt
	for c := range state {
		if c.x < minX {
			minX = c.x
		}
	}

	minY := maxInt
	for c := range state {
		if c.y < minY {
			minY = c.y
		}
	}

	minZ := maxInt
	for c := range state {
		if c.z < minZ {
			minZ = c.z
		}
	}

	minW := maxInt
	for c := range state {
		if c.w < minW {
			minW = c.w
		}
	}

	return coord4{
		x: minX,
		y: minY,
		z: minZ,
		w: minW,
	}
}

func maxPoint4(state map[coord4]struct{}) coord4 {
	minInt := -100000000000000

	maxX := minInt
	for c := range state {
		if c.x > maxX {
			maxX = c.x
		}
	}

	maxY := minInt
	for c := range state {
		if c.y > maxY {
			maxY = c.y
		}
	}

	maxZ := minInt
	for c := range state {
		if c.z > maxZ {
			maxZ = c.z
		}
	}

	maxW := minInt
	for c := range state {
		if c.w > maxW {
			maxW = c.w
		}
	}

	return coord4{
		x: maxX,
		y: maxY,
		z: maxZ,
		w: maxW,
	}
}
