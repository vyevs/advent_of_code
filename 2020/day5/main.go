package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()

	seats := convertToSeats(input)

	fmt.Println("Part 1:")
	doPart1(seats)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(seats)
}

type seat struct {
	row, col int
}

func (s seat) id() int {
	return s.row*8 + s.col
}

func convertToSeats(strs []string) []seat {
	seats := make([]seat, 0, len(strs))

	for _, str := range strs {
		seats = append(seats, toSeat(str))
	}

	return seats
}

func toSeat(bp string) seat {
	var seat seat
	{
		low, high := 0, 127
		for i := 0; i < 7; i++ {
			c := bp[i]
			if c == 'F' {
				high = (low + high) / 2
			} else if c == 'B' {
				low = (low + high + 1) / 2
			}
		}
		seat.row = low
	}

	{
		low, high := 0, 7
		for i := 7; i < 10; i++ {
			c := bp[i]
			if c == 'L' {
				high = (low + high) / 2
			} else if c == 'R' {
				low = (low + high + 1) / 2
			}
		}
		seat.col = low
	}

	return seat
}

func doPart1(seats []seat) {
	highestID := seats[0].id()
	for i := 1; i < len(seats); i++ {
		seatID := seats[i].id()
		if seatID > highestID {
			highestID = seatID
		}
	}

	fmt.Printf("Highest seat ID is %d\n", highestID)
}

func doPart2(seats []seat) {
	existingSeatIDs := make(map[int]struct{}, len(seats))

	for _, s := range seats {
		existingSeatIDs[s.id()] = struct{}{}
	}

	mySeatID := -1

	for r := 0; r < 128; r++ {
		for c := 0; c < 8; c++ {
			seatID := r*8 + c

			_, haveSeatID := existingSeatIDs[seatID]
			_, haveLowerID := existingSeatIDs[seatID-1]
			_, haveHigherID := existingSeatIDs[seatID+1]

			if !haveSeatID && haveLowerID && haveHigherID {
				mySeatID = seatID
				break
			}
		}
	}

	fmt.Printf("My seat ID is %d\n", mySeatID)
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	input := make([]string, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, line)
	}

	return input
}
