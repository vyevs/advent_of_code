package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	instrStrs := readInstructionStrs()

	instrs := parseInstructionStrs(instrStrs)

	fmt.Println("Part 1:")
	doPart1(instrs)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(instrs)
	fmt.Println()

}

func doPart1(instrs []instruction) {

	execCts := make([]int, len(instrs))

	var pc, acc int
	for {
		if execCts[pc] == 1 {
			break
		}
		execCts[pc]++

		instr := instrs[pc]

		switch instr.op {
		case instructionOpAcc:
			acc += instr.arg

		case instructionOpJmp:
			pc += instr.arg
			continue
		case instructionOpNop:
		}

		pc++
	}

	fmt.Printf("Acc has value %d before any instruction is executed a second time\n", acc)
}

func doPart2(instrs []instruction) {
	jmpAndNopIndices := findJmpsAndNops(instrs)

	fmt.Printf("There are %d jmps and nops to try\n", len(jmpAndNopIndices))

	reverseJmpOrNop := func(instr *instruction) {
		if instr.op == instructionOpJmp {
			instr.op = instructionOpNop
		} else if instr.op == instructionOpNop {
			instr.op = instructionOpJmp
		}
	}

	for _, idx := range jmpAndNopIndices {
		instr := &instrs[idx]
		reverseJmpOrNop(instr)

		if terminates, acc := doesProgramTerminate(instrs); terminates {
			fmt.Printf("After swapping instruction %d, the program terminates and acc ends with value %d\n", idx, acc)
			break
		}

		reverseJmpOrNop(instr)
	}

}

// returns whether the program terminates, and if it does, the accumulator value at the end
func doesProgramTerminate(instrs []instruction) (bool, int) {
	execCts := make([]int, len(instrs))

	var pc, acc int
	for {
		if pc >= len(instrs) {
			break
		}
		if execCts[pc] == 1 {
			return false, 0
		}
		execCts[pc]++

		instr := instrs[pc]

		switch instr.op {
		case instructionOpAcc:
			acc += instr.arg
		case instructionOpJmp:
			pc += instr.arg
			continue
		case instructionOpNop:
		}

		pc++
	}

	return true, acc
}

func findJmpsAndNops(instrs []instruction) []int {
	indices := make([]int, 0, len(instrs))

	for i, instr := range instrs {
		if instr.op == instructionOpJmp || instr.op == instructionOpNop {
			indices = append(indices, i)
		}
	}

	return indices
}

type instructionOp int

const (
	instructionOpAcc instructionOp = iota
	instructionOpJmp
	instructionOpNop
)

type instruction struct {
	op  instructionOp
	arg int
}

func parseInstructionStrs(instrStrs []string) []instruction {
	instrs := make([]instruction, 0, len(instrStrs))

	for _, instrStr := range instrStrs {
		instr := parseInstructionStr(instrStr)
		instrs = append(instrs, instr)
	}

	return instrs
}

func parseInstructionStr(instrStr string) instruction {
	var instr instruction

	tokens := strings.Split(instrStr, " ")

	{
		var op instructionOp
		switch tokens[0] {
		case "acc":
			op = instructionOpAcc
		case "jmp":
			op = instructionOpJmp
		case "nop":
			op = instructionOpNop
		}

		instr.op = op
	}

	{
		argStr := tokens[1][1:]
		arg, _ := strconv.Atoi(argStr)

		if tokens[1][0] == '-' {
			arg = -arg
		}
		instr.arg = arg
	}

	return instr
}

func readInstructionStrs() []string {
	scanner := bufio.NewScanner(os.Stdin)
	instrs := make([]string, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()

		instrs = append(instrs, line)
	}

	return instrs
}
