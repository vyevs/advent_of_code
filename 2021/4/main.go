package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vyevs/vtools"
)

func main() {
	calledNums, boards := getInput()
	doPart1(boards, calledNums)
	doPart2(boards, calledNums)
}

func doPart1(boards []board, calledNums []int) {
	winnerBoard, winningNum := determineWinner(boards, calledNums)
	fmt.Printf("Part 1: %d\n", winnerBoard.sumOfUnseen()*winningNum)
}

func doPart2(boards []board, calledNums []int) {
	lastToWin, winningNum := determineLastToWin(boards, calledNums)
	fmt.Printf("Part 2: %d\n", lastToWin.sumOfUnseen()*winningNum)
}

func determineLastToWin(bs []board, calledNums []int) (lastToWin board, winningNumber int) {
	completed := make([]bool, len(bs)) // completed[i] is whether the ith has won.
	nIncomplete := len(bs)             // Number of incomplete boards left.

	for _, cn := range calledNums {
		for i := range bs {
			b := &bs[i]

			if !completed[i] {
				b.markSeen(cn)

				if b.won() {
					if nIncomplete == 1 {
						return *b, cn
					} else {
						completed[i] = true
						nIncomplete--
					}
				}
			}
		}
	}

	panic("no winner")
}

func determineWinner(bs []board, calledNums []int) (winner board, winningNumber int) {
	for _, cn := range calledNums {
		for i := range bs {
			b := &bs[i]

			b.markSeen(cn)

			if b.won() {
				return *b, cn
			}
		}
	}

	panic("no winner")
}

type board struct {
	nums [5][5]int
	seen [5][5]bool
}

func (b *board) markSeen(called int) {
	for i, row := range b.nums {
		for j, v := range row {
			if v == called {
				b.seen[i][j] = true
			}
		}
	}
}

func (b board) won() bool {
	return b.completedAtLeastOneRow() || b.completedAtLeastOneColumn()
}

func (b board) completedAtLeastOneRow() bool {
outer:
	for r := range 5 {
		for c := range 5 {
			if !b.seen[r][c] {
				continue outer
			}
		}
		return true
	}
	return false
}

func (b board) completedAtLeastOneColumn() bool {
outer:
	for c := range 5 {
		for r := range 5 {
			if !b.seen[r][c] {
				continue outer
			}
		}
		return true
	}
	return false
}

func (b board) sumOfUnseen() int {
	var sum int
	for r := range 5 {
		for c := range 5 {
			if !b.seen[r][c] {
				sum += b.nums[r][c]
			}
		}
	}
	return sum
}

func (b board) String() string {
	var buf strings.Builder
	buf.Grow(128)
	for rowN, row := range b.nums {
		for i, v := range row {
			if i != 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(fmt.Sprintf("%2d", v))
		}
		if rowN < len(b.nums)-1 {
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func getInput() ([]int, []board) {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	line1 := lines[0]
	line1Split := strings.Split(line1, ",")
	calledNums := make([]int, 0, len(line1Split))
	for _, calledStr := range line1Split {
		calledNum, _ := strconv.Atoi(calledStr)
		calledNums = append(calledNums, calledNum)
	}

	boards := make([]board, 0, 64)
	lineNum := 1
	for lineNum < len(lines) { // for the empty line at the beginning
		var b board
		for i := range 5 {
			line := lines[lineNum]
			lineNum++

			nums := vtools.SplitWS(line)

			for j := range 5 {
				v, _ := strconv.Atoi(nums[j])
				b.nums[i][j] = v
			}
		}
		boards = append(boards, b)
	}

	return calledNums, boards
}
