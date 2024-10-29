package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/vyevs/vtools"
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	schematic := getEngineSchematic()
	fmt.Printf("%d lines\n", len(schematic))

	sum := sumPartNumbers(schematic)
	fmt.Printf("part 1: the sum of the engine part numbers is %d\n", sum)

	gearRatioSum := sumGearRatios(schematic)
	fmt.Printf("part 2: the sum of the gear ratios is %d\n", gearRatioSum)
}

func sumGearRatios(schematic [][]byte) int {
	var sum int

	nums := parseNums(schematic)
	slices.SortFunc(nums, func(a, b num) int {
		if a.gearLoc[0] < b.gearLoc[0] {
			return -1
		}
		if a.gearLoc[1] < b.gearLoc[1] {
			return -1
		}
		return 1
	})
	for i := range len(nums) - 1 {
		n1, n2 := nums[i], nums[i+1]
		if n1.gearLoc == n2.gearLoc {
			sum += n1.val * n2.val
		}

	}

	return sum
}

func parseNums(schematic [][]byte) []num {
	gearR, gearC := -1, -1
	var arr [4]byte
	buf := arr[:0]
	nums := make([]num, 0, 1024)
	useBuf := func() {
		if len(buf) > 0 {
			if gearR != -1 {
				partNum := atoi(buf)
				n := num{
					val:     partNum,
					gearLoc: [2]int{gearR, gearC},
				}
				nums = append(nums, n)
				gearR, gearC = -1, -1
			}
			buf = buf[:0]
		}
	}

	for r, row := range schematic {
		for c, char := range row {
			if isNum(char) {
				if gearR == -1 {
					gearR, gearC = adjacentGear(schematic, r, c)
				}

				buf = append(buf, char)
			} else {
				useBuf()
			}
		}
	}

	useBuf()

	return nums
}

type num struct {
	val     int
	gearLoc [2]int
}

func sumPartNumbers(schematic [][]byte) int {
	var sum int
	var seenSymbol bool
	var arr [4]byte
	buf := arr[:0]
	useBuf := func() {
		if len(buf) > 0 {
			if seenSymbol {
				partNum := atoi(buf)
				sum += partNum
				seenSymbol = false
			}
			buf = buf[:0]
		}
	}

	for r, row := range schematic {
		for c, char := range row {
			if isNum(char) {
				if !seenSymbol && adjacentToSymbol(schematic, r, c) {
					seenSymbol = true
				}
				buf = append(buf, char)
			} else {
				useBuf()
			}
		}
	}

	useBuf()

	return sum
}

func adjacentToSymbol(grid [][]byte, r, c int) bool {
	isSymbol := func(r, c int) bool {
		if r < 0 || r > len(grid)-1 || c < 0 || c > len(grid[r])-1 {
			return false
		}
		v := grid[r][c]
		return v != '.' && !isNum(v)
	}

	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, d := range directions {
		dr, dc := d[0], d[1]

		if isSymbol(r+dr, c+dc) {
			return true
		}
	}
	return false
}

func adjacentGear(grid [][]byte, r, c int) (int, int) {
	isGear := func(r, c int) bool {
		if r < 0 || r > len(grid)-1 || c < 0 || c > len(grid[r])-1 {
			return false
		}
		v := grid[r][c]
		return v == '*'
	}

	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, d := range directions {
		dr, dc := d[0], d[1]

		if isGear(r+dr, c+dc) {
			return r + dr, c + dc
		}
	}
	return -1, -1
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func getEngineSchematic() [][]byte {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	grid := make([][]byte, 0, len(lines))
	for _, line := range lines {
		row := make([]byte, 0, len(line))
		for c := range vtools.StrBytes(line) {
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	return grid
}

func atoi(bs []byte) int {
	v, err := strconv.Atoi(string(bs))
	if err != nil {
		panic(err)
	}
	return v
}
