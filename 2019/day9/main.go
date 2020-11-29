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
    	// opcode 1
	// 3 operands
	// adds the 2 operands and stores the result at position 3rd operand
	instructionTypeAdd instructionType = 1 + iota

	// opcode 2
	// 3 operands
	// multiplies 2 operands and stores result at position 3rd operand
	instructionTypeMultiply

	// opcode 3
	// 1 operand
	// takes an integer input and stores it at position 1st operand
	instructionTypeInput

	// opcode 4
	// 1 operand
	// outputs value at 1st operand
	instructionTypeOutput

	// opcode 5
	// 2 operands
	// if operand 1 is non-zero, jumps program to pc of value of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfTrue

	// opcode 6
	// 2 operands
	// if operand 1 is 0, jumps program to pc of value of of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfFalse

	// opcode 7
	// 3 operands
	// if operand 1 is less than operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeLessThan

	// opcode 8
	// 3 operands
	// if openrand 1 equals operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeEquals

	// opcode 9
	// 1 operand
	// modifies the relative base for the computer by the only operand provided, increasing it if it is positive
	// decreasing it if it is negative
	instructionTypeAdjustRelativeBase

	// opcode 99
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
	// the operand will be interpreted as an offset from the computer's relative base
	opModeRelative
)

type instruction struct {
	typ instructionType

	op1, op2, op3    int
	op1Mode, op2Mode opMode
}

type computer struct {
	instrs          []int
	pc              int // program counter
	getInput func() int
	relativeBase int
	running         bool
}

func (c computer) parseInstructionInt(instr int) instruction {
	if c.instrs[c.pc] != instr {
		panic(fmt.Sprintf("instruction %d at pc %d does not equal provided instruction %d", c.instrs[c.pc], c.pc, instr))
	}

	var actualInstr instruction

	// the instruction opcode consists of only the ones and tens digits, all other digits are irrelevant
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

	} else if typ == instructionTypeAdjustRelativeBase {
	    	actualInstr.op1Mode = opMode((instr / 100) % 10)

		actualInstr.op1 = c.instrs[c.pc+1]

	} else if typ == instructionTypeHalt {

	} else {
		panic(fmt.Sprintf("encountered unknown instruction type %d", typ))
	}

	return actualInstr
}

// returnrs whether the computer can continue running or if it has reached the halt instruction
func (c *computer) run() {
	c.running = true


	for c.running {
		instrInt := c.instrs[c.pc]

		instr := c.parseInstructionInt(instrInt)

		c.executeInstruction(instr)
	}

	log.Print("computer stopped")
}

// returns whether we needed input to execute an instruction but there was no input ready
func (c *computer) executeInstruction(instr instruction) {
	switch instr.typ {
	case instructionTypeAdd:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		c.setValueAtAddress(instr.op3, v1+v2)

		c.pc += 4

	case instructionTypeMultiply:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		c.setValueAtAddress(instr.op3, v1 * v2)

		c.pc += 4

	case instructionTypeInput:
	    	input := c.getInput()
		c.setValueAtAddress(instr.op1, input)
		c.pc += 2

	case instructionTypeOutput:
		var output int
		if instr.op1Mode == opModeImmediate {
			output = instr.op1
		} else if instr.op1Mode == opModePosition {
			output = c.getValueAtAddress(instr.op1)
		} else if instr.op1Mode == opModeRelative {
			output = c.getValueAtAddress(c.relativeBase + instr.op1)
		} else {
			panic(fmt.Sprintf("unknown op mode %d", instr.op1Mode))
		}


		log.Printf("output: %d", output)

		c.pc += 2

	case instructionTypeJumpIfTrue:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		if v1 != 0 {
			c.pc = v2
		} else {
			c.pc += 3
		}

	case instructionTypeJumpIfFalse:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		if v1 == 0 {
			c.pc = v2
		} else {
			c.pc += 3
		}

	case instructionTypeLessThan:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		if v1 < v2 {
			c.setValueAtAddress(instr.op3, 1)
		} else {
			c.setValueAtAddress(instr.op3, 0)
		}

		c.pc += 4

	case instructionTypeEquals:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		if v1 == v2 {
			c.setValueAtAddress(instr.op3, 1)
		} else {
			c.setValueAtAddress(instr.op3, 0)
		}

		c.pc += 4
	
	case instructionTypeAdjustRelativeBase:
	    	adjustBy := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		c.relativeBase += adjustBy

		c.pc += 2

	case instructionTypeHalt:
		c.running = false

	default:
		panic(fmt.Sprintf("unknown instruction type %d", instr.typ))
	}
}

func (c *computer) loadValueBasedOnOpMode(operand int, operandMode opMode) int {
    	switch operandMode {
	    case opModeImmediate:
		return operand

	    case opModePosition:
		return c.getValueAtAddress(operand)

	    case opModeRelative:
		return c.getValueAtAddress(c.relativeBase + operand)
	}
	panic(fmt.Sprintf("unknown op mode %d", operandMode))
}

// this function retrieves the value located in the computer's memory at address a
// this is necessary because our computer must have memory beyond the initial program
// the memory of the computer is grown everytime we try to access a value that is currently out of bounds
// memory that has not been accessed yet is zeroed
func (c *computer) getValueAtAddress(a int) int {
    if a >= len(c.instrs) {
	c.growMemory(a+1) // we must grow to a+1 because that would make a the last valid index
    }
    return c.instrs[a]
}

// this function sets the value located in the computer's memory at address a to value v
// this is necessary because our computer must have memory beyond the initial program
// the memory of the computer is grown everytime we try to access a value that is currently out of bounds
// memory that has not been accessed yet is zeroed
func (c *computer) setValueAtAddress(a, v int) {
    if a >= len(c.instrs) {
	c.growMemory(a+1) // we must grow to a+1 because that would make a the last valid index
    }
    c.instrs[a] = v
}

// grows the computer's memory to size n, so that the last valid index is n-1
func (c *computer) growMemory(n int) {
    newMemory := make([]int, n)
    copy(newMemory, c.instrs)
    c.instrs = newMemory
}

func main() {
	defer func(start time.Time) {
		log.Printf("that took %s", time.Since(start))
	}(time.Now())

	instrs := readInputInstructions()

	doPart1(instrs)
}

func doPart1(instrs []int) {
    c := computer {
	instrs: instrs,
	getInput: func() int {
	    return 1
	},
    }

    c.run()
}

func readInputInstructions() []int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	instrStrs := strings.Split(line, ",")

	instrs := make([]int, 0, len(instrStrs))
	for _, str := range instrStrs {
		instr, _ := strconv.Atoi(str)
		instrs = append(instrs, instr)
	}

	return instrs
}
