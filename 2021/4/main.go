package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calledNums, boards := getInput()
	doPart1(boards, calledNums)
	doPart2(boards, calledNums)
}

func doPart1(boards []board, calledNums []int) {
	winnerBoard, winningNum := determineWinner(boards, calledNums)
	fmt.Printf("Part 1: %d\n", winnerBoard.sumOfUnseen() * winningNum)
}

func doPart2(boards []board, calledNums []int) {
	lastToWin, winningNum := determineLastToWin(boards, calledNums)
	fmt.Printf("Part 2: %d\n", lastToWin.sumOfUnseen() * winningNum)
}

func determineLastToWin(bs []board, calledNums []int) (lastToWin board, winningNumber int) {
	incompleteBoards := make(map[int]struct{}, len(bs))
	for i := range bs {
		incompleteBoards[i] = struct{}{}
	}
	
	for _, cn := range calledNums {
		for i := range bs {
			b := &bs[i]
			
			if _, isIncomplete := incompleteBoards[i]; isIncomplete {
				b.markSeen(cn)
				
				if b.won() {
					if len(incompleteBoards) == 1 {
						return *b, cn
					} else {
						delete(incompleteBoards, i)
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
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
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
	for c := 0; c < 5; c++ {
		for r := 0; r < 5; r++ {
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
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
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
				buf.WriteRune(' ')
			}
			buf.WriteString(fmt.Sprintf("%2d", v))
		}
		if rowN < len(b.nums) - 1 {
			buf.WriteRune('\n')
		}
	}
	return buf.String()
}

func getInput() ([]int, []board) {
	scnr := bufio.NewScanner(os.Stdin)

	scnr.Scan()
	line1 := scnr.Text()
	line1Split := strings.Split(line1, ",")
	calledNums := make([]int, 0, len(line1Split))
	for _, calledStr := range line1Split {
		calledNum, _ := strconv.Atoi(calledStr)
		calledNums = append(calledNums, calledNum)
	}

	boards := make([]board, 0, 64)
	for scnr.Scan() { // for the empty line at the beginning
		var b board
		for i := 0; i < 5; i++ {
			scnr.Scan()
			line := scnr.Text()
			nums := strings.Split(line, " ")
			nums = removeEmpties(nums)
			
			for j := 0; j < 5; j++ {
				v, _ := strconv.Atoi(nums[j])
				b.nums[i][j] = v
			}
		}
		boards = append(boards, b)
	}

	return calledNums, boards
}

func removeEmpties(strs []string) []string {
	ret := make([]string, 0, 5)
	for _, s := range strs {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return ret
}