package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println("Part1:")
	doPart1(input)
	fmt.Println()

	fmt.Println("Part2:")
	doPart2(input)
}

func doPart1(bags bags) {
	var ct int
	for _, bag := range bags.bags {
		if bag.color != "shiny gold" && bagContainsShinyGold(bag, bags) {
			ct++
		}
	}

	fmt.Printf("%d bags can eventually contain at least one shiny gold bag", ct)
}

func bagContainsShinyGold(b bag, bags bags) bool {
	for _, color := range b.mustContainColors {
		if color == "shiny gold" {
			return true
		}

		childBag := bags.colorToBag[color]

		if bagContainsShinyGold(childBag, bags) {
			return true
		}
	}

	return false
}

func doPart2(bags bags) {
	goldBag := bags.colorToBag["shiny gold"]

	var totalCt int
	for i, childBagColor := range goldBag.mustContainColors {
		childBag := bags.colorToBag[childBagColor]

		totalCt += goldBag.mustContainCts[i] * nBagsInsideIncludingTopBag(childBag, bags)
	}

	fmt.Printf("%d bags are required inside the shiny gold bag\n", totalCt)
}

func nBagsInsideIncludingTopBag(b bag, bags bags) int {
	totalCt := 1

	for i, ct := range b.mustContainCts {
		innerBagColor := b.mustContainColors[i]

		childBag := bags.colorToBag[innerBagColor]

		totalCt += ct * nBagsInsideIncludingTopBag(childBag, bags)
	}

	return totalCt
}

type bags struct {
	bags       []bag
	colorToBag map[string]bag
}

type bag struct {
	color string

	mustContainColors []string
	mustContainCts    []int
}

func readInput() bags {
	scanner := bufio.NewScanner(os.Stdin)

	var bags bags
	bags.bags = make([]bag, 0, 1024)
	bags.colorToBag = make(map[string]bag, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		bagColor := tokens[0] + " " + tokens[1]

		var bag bag

		bag.color = bagColor

		if tokens[4] == "no" {
			continue
		}
		for i := 4; i < len(tokens); i += 4 {
			reqCtStr := tokens[i]
			reqCt, _ := strconv.Atoi(reqCtStr)

			reqColor := tokens[i+1] + " " + tokens[i+2]

			bag.mustContainColors = append(bag.mustContainColors, reqColor)
			bag.mustContainCts = append(bag.mustContainCts, reqCt)
		}

		bags.bags = append(bags.bags, bag)
		bags.colorToBag[bagColor] = bag
	}

	return bags
}
