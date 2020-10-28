package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	masses := readMasses()

	//solvePart1(masses)
	solvePart2(masses)
}

func readMasses() []int {
	scanner := bufio.NewScanner(os.Stdin)

	masses := make([]int, 0, 64)
	for scanner.Scan() {
		line := scanner.Text()

		mass, _ := strconv.Atoi(line)

		masses = append(masses, mass)
	}

	return masses
}

func solvePart1(masses []int) {
	var fuelRequired int
	for _, mass := range masses {
		fuelRequired += mass/3 - 2
	}

	fmt.Printf("part 1: %v", fuelRequired)
}

func solvePart2(masses []int) {

	var fuelRequired int
	for _, mass := range masses {
		initialFuelForMass := mass/3 - 2
		fuelForFuelOfMass := calculateFuelRequiredForFuel(initialFuelForMass)

		fuelRequired += initialFuelForMass + fuelForFuelOfMass
	}

	fmt.Printf("part 1: %v", fuelRequired)
}

func calculateFuelRequiredForFuel(massOfFuel int) int {
	var totalExtraFuel int
	for extraFuel := massOfFuel/3 - 2; extraFuel > 0; extraFuel = extraFuel/3 - 2 {
		totalExtraFuel += extraFuel
	}
	return totalExtraFuel
}
