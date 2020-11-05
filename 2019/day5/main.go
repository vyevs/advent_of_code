package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type instructionType int

const (
	// 3 operands
	// adds the 2 operands and stores the result at position 3rd operand
	instructionTypeAdd instructionType = 1 + iota

	// 3 operands
	// multiplies 2 operands and stores result at position 3rd operand
	instructionTypeMultiply

	// 1 operand
	// takes an integer input and stores it at position 1st operand
	instructionTypeInput

	// 1 operand
	// outputs value at 1st operand
	instructionTypeOutput

	// 2 operands
	// if operand 1 is non-zero, jumps program to pc of value of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfTrue

	// 2 operands
	// if operand 1 is 0, jumps program to pc of value of of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfFalse

	// 3 operands
	// if operand 1 is less than operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeLessThan

	// 3 operands
	// if openrand 1 equals operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeEquals

	// no operands
	// halts the computer
	instructionTypeHalt instructionType = 99
)

type opMode int

const (
	// the operand will be be interpreted as the position of the value to be operated on
	opModePosition opMode = iota
	// the operand will be interpreted as an immediate value
	opModeImmediate
)

type instruction struct {
	typ instructionType

	op1, op2, op3    int
	op1Mode, op2Mode opMode
}

type computer struct {
	getInput      func() int
	produceOutput func(int)
	instrs        []int
	pc            int // program counter
	running       bool
}

func (c computer) parseInstructionInt(instr int) instruction {
	if c.instrs[c.pc] != instr {
		panic(fmt.Sprintf("instruction %d at pc %d does not equal provided instruction %d", c.instrs[c.pc], c.pc, instr))
	}

	var actualInstr instruction

	typ := instructionType(instr % 100)
	actualInstr.typ = typ

	if typ == instructionTypeAdd || typ == instructionTypeMultiply ||
		typ == instructionTypeLessThan || typ == instructionTypeEquals {

		actualInstr.op1Mode = opMode((instr / 100) % 10)
		actualInstr.op2Mode = opMode((instr / 1000) % 10)

		actualInstr.op1 = c.instrs[c.pc+1]
		actualInstr.op2 = c.instrs[c.pc+2]
		actualInstr.op3 = c.instrs[c.pc+3]

	} else if typ == instructionTypeInput {
		actualInstr.op1 = c.instrs[c.pc+1]

	} else if typ == instructionTypeOutput {
		actualInstr.op1Mode = opMode((instr / 100) % 10)

		actualInstr.op1 = c.instrs[c.pc+1]

	} else if typ == instructionTypeJumpIfTrue || typ == instructionTypeJumpIfFalse {
		actualInstr.op1Mode = opMode((instr / 100) % 10)
		actualInstr.op2Mode = opMode((instr / 1000) % 10)

		actualInstr.op1 = c.instrs[c.pc+1]
		actualInstr.op2 = c.instrs[c.pc+2]

	} else if typ == instructionTypeHalt {

	} else {
		panic(fmt.Sprintf("unknown instruction type %d", typ))
	}

	return actualInstr
}

func (c *computer) run(instrs []int) {
	c.running = true
	c.instrs = instrs

	for c.running {
		instrInt := instrs[c.pc]

		instr := c.parseInstructionInt(instrInt)

		c.executeInstruction(instr)
	}
}

func (c *computer) executeInstruction(instr instruction) {
	switch instr.typ {
	case instructionTypeAdd:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		c.instrs[instr.op3] = v1 + v2

		c.pc += 4

	case instructionTypeMultiply:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		c.instrs[instr.op3] = v1 * v2

		c.pc += 4

	case instructionTypeInput:
		input := c.getInput()
		c.instrs[instr.op1] = input

		c.pc += 2

	case instructionTypeOutput:
		var output int
		if instr.op1Mode == opModeImmediate {
			output = instr.op1
		} else if instr.op1Mode == opModePosition {
			output = c.instrs[instr.op1]
		} else {
			panic(fmt.Sprintf("unknown op mode %d", instr.op1Mode))
		}

		c.produceOutput(output)

		c.pc += 2

	case instructionTypeJumpIfTrue:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		if v1 != 0 {
			c.pc = v2
		} else {
			c.pc += 3
		}

	case instructionTypeJumpIfFalse:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		if v1 == 0 {
			c.pc = v2
		} else {
			c.pc += 3
		}

	case instructionTypeLessThan:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		if v1 < v2 {
			c.instrs[instr.op3] = 1
		} else {
			c.instrs[instr.op3] = 0
		}

		c.pc += 4

	case instructionTypeEquals:
		v1 := loadValueBasedOnOpMode(c.instrs, instr.op1, instr.op1Mode)
		v2 := loadValueBasedOnOpMode(c.instrs, instr.op2, instr.op2Mode)

		if v1 == v2 {
			c.instrs[instr.op3] = 1
		} else {
			c.instrs[instr.op3] = 0
		}

		c.pc += 4

	case instructionTypeHalt:
		c.running = false

	default:
		panic(fmt.Sprintf("unknown instruction type %d", instr.typ))
	}
}

func loadValueBasedOnOpMode(instrs []int, op int, opMode opMode) int {
	if opMode == opModeImmediate {
		return op
	} else if opMode == opModePosition {
		return instrs[op]
	}
	panic(fmt.Sprintf("unknown op mode %d", opMode))
}

func main() {
	defer func(start time.Time) {
		log.Printf("that took %s", time.Since(start))
	}(time.Now())
	doIt()
}

func doIt() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	instrStrs := strings.Split(line, ",")

	instrs := make([]int, 0, len(instrStrs))
	for _, str := range instrStrs {
		instr, _ := strconv.Atoi(str)
		instrs = append(instrs, instr)
	}

	computer := computer{
		getInput: func() int {
			log.Print("computer got input")
			return 5
		},
		produceOutput: func(output int) {
			log.Printf("computer produced output %d", output)
		},
	}

	computer.run(instrs)
}
