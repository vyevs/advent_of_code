package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	earliestDepart, busIDs := readInput()

	fmt.Println("Part 1:")
	doPart1(earliestDepart, busIDs)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(busIDs)
}

func doPart1(earliestDepart int, busIDs []int) {

	var busToTake int
	departAt := earliestDepart * 2
	for _, busID := range busIDs {
		if busID == 0 {
			continue
		}
		busDepartsAt := earliestDepart + (busID - (earliestDepart % busID))

		if busDepartsAt < departAt {
			departAt = busDepartsAt
			busToTake = busID
		}
	}

	waitTime := departAt - earliestDepart
	fmt.Printf("You should take bus ID %d, it will depart at %d (you wait %d minutes)\n", busToTake, departAt, waitTime)
	fmt.Printf("%d * %d = %d\n", busToTake, waitTime, busToTake*waitTime)
}

func doPart2(busIDs []int) {
	time := busIDs[0]
	add := 1
	for i := 0; i < len(busIDs); i++ {
		busID := busIDs[i]
		if busID == 0 {
			continue
		}

		for (time+i)%busID != 0 {
			time += add
		}
		add *= busID
	}

	fmt.Printf("%d\n", time)
}

func readInput() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	earliestDepart, _ := strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()
	busIDStrs := strings.Split(line, ",")
	busIDs := make([]int, 0, len(busIDStrs))
	for _, busIDStr := range busIDStrs {
		if busIDStr == "x" {
			busIDs = append(busIDs, 0)
		} else {
			busID, _ := strconv.Atoi(busIDStr)
			busIDs = append(busIDs, busID)
		}
	}

	return earliestDepart, busIDs
}
