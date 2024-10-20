package main

import (
	"fmt"
	"io"
	"os"
	
	"golang.org/x/exp/maps"
)

func main() {
	moves := getInput(os.Stdin)
	part1(moves)
}

func part1(ms []move) {

	var h, t v2
	unitMoves := getUnitTailMoves()
	visited := make(map[v2]struct{}, 1>>16)
	for _, m := range ms {
		hMove := m.dir.toV2()
		for i := 0; i < m.mag; i++ {
			//printGrid(h, t)
	
			visited[t] = struct{}{}
			
			h = h.add(hMove)

			delta := h.sub(t)
			tMove := unitMoves[delta]
			
			t = t.add(tMove)
		}
	}
	
	//printFinalGrid(visited)

	fmt.Printf("tail visited %d locations\n", len(visited))
}


func printGrid(h, t v2) {
	topLeft := v2{x: min(h.x, t.x, 0) - 1, y: max(h.y, t.y, 0) + 1}
	bottomRight := v2{x: max(h.x, t.x, 0) + 1, y: min(h.y, t.y, 0) - 1}
	
	for i := topLeft.y; i >= bottomRight.y; i-- { // for each row in the grid
		for j := topLeft.x; j <= bottomRight.x; j++ { // for each column in the grid
			
			if i == h.y && j == h.x {
				fmt.Print("H")
			} else if i == t.y && j == t.x {
				fmt.Print("T")
			} else if i == 0 && j == 0 { // the start marker
				fmt.Print("s")
			} else {
				fmt.Print(".")
			}
			
		}
		fmt.Println()
	}
	fmt.Println()
}

func printFinalGrid(visited map[v2]struct{}) {
	ks := maps.Keys(visited)
	
	xs := make([]int, 0, len(ks))
	ys := make([]int, 0, len(ks))
	for _, k := range ks {
		xs = append(xs, k.x)
		ys = append(ys, k.y)
	}
	
	topLeft := v2{x: min(xs...) - 1, y: max(ys...) + 1}
	bottomRight := v2{x: max(xs...) + 1, y: min(ys...) -1}
	
	for i := topLeft.y; i >= bottomRight.y; i-- { // for each row in the grid
		for j := topLeft.x; j <= bottomRight.x; j++ { // for each column in the grid
			
			if i == 0 && j == 0 {
				fmt.Print("s")
			} else if _, didVisit := visited[v2{x: j, y: i}]; didVisit {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			
		}
		fmt.Println()
	}
	fmt.Println()
}

func getUnitTailMoves() map[v2]v2 {
	moves := make(map[v2]v2, 20)
	
	// Head is immediately next to the tail, either vertically, horizontally or diagonally. Or they're on top of each other.
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			v := v2{x: dx, y: dy}
			moves[v] = v2{x:0, y:0}
		}
	}
	
	addMove := func(dx, dy, dtx, dty int) {
		moves[v2{x: dx, y: dy}] = v2{x: dtx, y: dty}
	}

	// Horizontal moves
	addMove(2, 0, 1, 0)
	addMove(-2, 0, -1, 0)
	
	// Vertical moves
	addMove(0, 2, 0, 1)
	addMove(0, -2, 0, -1)

	// Diagonal moves
	addMove(1, 2, 1, 1)
	addMove(2, 1, 1, 1)
	addMove(2, -1, 1, -1)
	addMove(1, -2, 1, -1)
	addMove(-1, 2, -1, 1)
	addMove(-2, 1, -1, 1)
	addMove(-1, -2, -1, -1)
	addMove(-2, -1, -1, -1)
	
	fmt.Printf("there are %d move mappings\n", len(moves))
	
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(vs ...int) int {
	m := vs[0]
	for _, v := range vs {
		if v < m {
			m = v
		}
	}
	return m
}

func max(vs ...int) int {
	m := vs[0]
	for _, v := range vs {
		if v > m {
			m = v
		}
	}
	return m
}

type v2 struct {
	x, y int
}

func (v v2) String() string {
	return fmt.Sprintf("[%d, %d]", v.x, v.y)
}

func (v v2) add(o v2) v2 {
	return v2{
		x: v.x + o.x,
		y: v.y + o.y,
	}
}

func (v v2) sub(o v2) v2 {
	return v2 {
		x: v.x - o.x,
		y: v.y - o.y,
	}
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) toV2() v2 {
	switch d {
	case up:
		return v2{x: 0, y: 1}
	case down:
		return v2{x: 0, y: -1}
	case left:
		return v2{x: -1, y: 0}
	case right:
		return v2{x: 1, y: 0}
	}
	panic("")
}

type move struct {
	dir direction
	mag int
}

func (m move) toV2() v2 {
	switch m.dir {
	case up:
		return v2{x: 0, y: m.mag}
	case down:
		return v2{x: 0, y: -m.mag}
	case left:
		return v2{x: -m.mag, y: 0}
	case right:
		return v2{x: m.mag, y: 0}
	}
	panic("")
}

func strToDir(s string) direction {
	switch s {
	case "U":
		return up
	case "D":
		return down
	case "L":
		return left
	case "R":
		return right
	}
	panic("")
}

func getInput(r io.Reader) []move {
	ms := make([]move, 0, 1024)
	for {
		var dirStr string
		var mag int

		_, err := fmt.Fscan(r, &dirStr, &mag)
		if err != nil {
			break
		}

		m := move{
			dir: strToDir(dirStr),
			mag: mag,
		}
		ms = append(ms, m)
	}
	return ms
}
