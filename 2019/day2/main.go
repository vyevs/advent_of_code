package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	instrs := readIntegerInstructions()
	fmt.Println(instrs)

	// doPart1(instrs)
	doPart2(instrs)
}

func doPart1(instrs []int) {
	modifyInstrsForPart1(instrs)
	doInstrsPart1(instrs)

	fmt.Println(instrs)
}

func doPart2(instrs []int) {
	originalInstrs := make([]int, len(instrs))
	copy(originalInstrs, instrs)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(instrs, originalInstrs)

			instrs[1] = noun
			instrs[2] = verb

			doInstrsPart1(instrs)

			if instrs[0] == 19690720 {
				fmt.Printf("noun = %d, verb = %d, 100 * noun + verb = %d", noun, verb, 100*noun+verb)
				return
			}
		}
	}

	panic("no answer found for part 2")
}

func doInstrsPart1(instrs []int) {

	var pc int
	for {
		opcode := instrs[pc]

		if opcode != 1 && opcode != 2 && opcode != 99 {
			panic(fmt.Sprintf("ran into invalid opcode %d at pc %d", opcode, pc))
		}

		if opcode == 99 {
			break
		}

		if opcode == 1 || opcode == 2 {
			srcPos1 := instrs[pc+1]
			srcPos2 := instrs[pc+2]
			destPos := instrs[pc+3]

			srcVal1 := instrs[srcPos1]
			srcVal2 := instrs[srcPos2]

			var resultVal int
			if opcode == 1 {
				resultVal = srcVal1 + srcVal2
			} else {
				resultVal = srcVal1 * srcVal2
			}

			instrs[destPos] = resultVal
		}

		pc += 4
	}
}

func readIntegerInstructions() []int {
	scanner := bufio.NewScanner(os.Stdin)

	instrs := make([]int, 0, 64)

	scanner.Scan()
	line := scanner.Text()

	iStrs := strings.Split(line, ",")
	for _, iStr := range iStrs {
		i, _ := strconv.Atoi(iStr)

		instrs = append(instrs, i)
	}
	return instrs
}

func modifyInstrsForPart1(instrs []int) {
	instrs[1] = 12
	instrs[2] = 2
}
