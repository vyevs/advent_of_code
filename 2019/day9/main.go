package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	//instructionWriter = os.Stdout
	instructionWriter = ioutil.Discard
	//memoryAccessWriter = os.Stdout
	memoryAccessWriter = ioutil.Discard
	//computerStateWriter = os.Stdout
	computerStateWriter = ioutil.Discard
)

type instructionType int

const (
	// opcode 1
	// 3 operands
	// adds the 2 operands and stores the result at position 3rd operand
	instructionTypeAdd instructionType = 1

	// opcode 2
	// 3 operands
	// multiplies 2 operands and stores result at position 3rd operand
	instructionTypeMultiply instructionType = 2

	// opcode 3
	// 1 operand
	// takes an integer input and stores it at position 1st operand
	instructionTypeInput instructionType = 3

	// opcode 4
	// 1 operand
	// outputs value at 1st operand
	instructionTypeOutput instructionType = 4

	// opcode 5
	// 2 operands
	// if operand 1 is non-zero, jumps program to pc of value of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfTrue instructionType = 5

	// opcode 6
	// 2 operands
	// if operand 1 is 0, jumps program to pc of value of of 2nd operand, otherwise goes to next instruction
	instructionTypeJumpIfFalse instructionType = 6

	// opcode 7
	// 3 operands
	// if operand 1 is less than operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeLessThan instructionType = 7

	// opcode 8
	// 3 operands
	// if operand 1 equals operand 2, stores 1 in the position of value of 3rd operand, otherwise 0
	instructionTypeEquals instructionType = 8

	// opcode 9
	// 1 operand
	// modifies the relative base for the computer by the only operand provided, increasing it if it is positive
	// decreasing it if it is negative
	instructionTypeAdjustRelativeBase instructionType = 9

	// opcode 99
	// no operands
	// halts the computer
	instructionTypeHalt instructionType = 99
)

var instructionTypeStrings = map[instructionType]string{
	instructionTypeAdd:                "ADD",
	instructionTypeMultiply:           "MUL",
	instructionTypeInput:              "IN",
	instructionTypeOutput:             "OUT",
	instructionTypeJumpIfTrue:         "JIT",
	instructionTypeJumpIfFalse:        "JIF",
	instructionTypeLessThan:           "LT",
	instructionTypeEquals:             "EQ",
	instructionTypeAdjustRelativeBase: "ARB",
	instructionTypeHalt:               "HALT",
}

func (t instructionType) String() string {
	str, isValid := instructionTypeStrings[t]

	if !isValid {
		panic(fmt.Sprintf("instructionType.String() invalid instructionType %d", t))
	}

	return fmt.Sprintf("%-3s", str)
}

func (t instructionType) usesOp1() bool {
	return t == instructionTypeAdd ||
		t == instructionTypeMultiply ||
		t == instructionTypeInput ||
		t == instructionTypeOutput ||
		t == instructionTypeJumpIfTrue ||
		t == instructionTypeJumpIfFalse ||
		t == instructionTypeLessThan ||
		t == instructionTypeEquals ||
		t == instructionTypeAdjustRelativeBase
}

func (t instructionType) usesOp2() bool {
	return t == instructionTypeAdd ||
		t == instructionTypeMultiply ||
		t == instructionTypeJumpIfTrue ||
		t == instructionTypeJumpIfFalse ||
		t == instructionTypeLessThan ||
		t == instructionTypeEquals
}

func (t instructionType) usesOp3() bool {
	return t == instructionTypeAdd ||
		t == instructionTypeMultiply ||
		t == instructionTypeLessThan ||
		t == instructionTypeEquals
}

type opMode int

const (
	// the operand will be be interpreted as the position of the value to be operated on
	opModePosition opMode = 0
	// the operand will be interpreted as an immediate value
	opModeImmediate opMode = 1
	// the operand will be interpreted as an offset from the computer's relative base
	opModeRelative opMode = 2
)

type instruction struct {
	typ instructionType

	op1, op2, op3             int
	op1Mode, op2Mode, op3Mode opMode
}

func (i instruction) String() string {
	var buf strings.Builder
	buf.Grow(64)

	buf.WriteString(i.typ.String())
	buf.WriteRune(' ')

	if i.typ.usesOp1() {
		buf.WriteRune(' ')

		if i.op1Mode == opModePosition || i.op1Mode == opModeRelative {
			buf.WriteRune('[')
		}

		if i.op1Mode == opModeRelative {
			buf.WriteString("rb")
			if i.op1 >= 0 {
				buf.WriteRune('+')
			}
		}

		buf.WriteString(strconv.Itoa(i.op1))

		if i.op1Mode == opModePosition || i.op1Mode == opModeRelative {
			buf.WriteRune(']')
		}
	}

	if i.typ.usesOp2() {
		buf.WriteRune(',')

		if i.op2Mode == opModePosition || i.op2Mode == opModeRelative {
			buf.WriteRune('[')
		}

		if i.op2Mode == opModeRelative {
			buf.WriteString("rb")
			if i.op2 > 0 {
				buf.WriteRune('+')
			}
		}

		buf.WriteString(strconv.Itoa(i.op2))

		if i.op2Mode == opModePosition || i.op2Mode == opModeRelative {
			buf.WriteRune(']')
		}
	}

	if i.typ.usesOp3() {
		buf.WriteRune(',')
		buf.WriteRune('[')
		if i.op3Mode == opModeRelative {
			buf.WriteString("rb")
			if i.op2 > 0 {
				buf.WriteRune('+')
			}
		}
		buf.WriteString(strconv.Itoa(i.op3))
		buf.WriteRune(']')

	}

	return buf.String()
}

type computer struct {
	memory        []int
	pc            int // program counter
	getInput      func() int
	produceOutput func(v int)
	relativeBase  int
	running       bool
}

func (c computer) parseInstructionInt(instr int) instruction {
	if c.memory[c.pc] != instr {
		panic(fmt.Sprintf("instruction %d at pc %d does not equal provided instruction %d", c.memory[c.pc], c.pc, instr))
	}

	var actualInstr instruction

	// the instruction opcode consists of only the ones and tens digits, all other digits are irrelevant
	typ := instructionType(instr % 100)
	actualInstr.typ = typ

	if typ == instructionTypeAdd || typ == instructionTypeMultiply ||
		typ == instructionTypeLessThan || typ == instructionTypeEquals {

		actualInstr.op1Mode = op1Mode(instr)
		actualInstr.op2Mode = op2Mode(instr)
		actualInstr.op3Mode = op3Mode(instr)

		if actualInstr.op3Mode == opModeImmediate {
			panic(fmt.Sprintf("got opModeImmediate for instruction that writes to address of 3rd operand, instr: %d", instr))
		}

		actualInstr.op1 = c.memory[c.pc+1]
		actualInstr.op2 = c.memory[c.pc+2]
		actualInstr.op3 = c.memory[c.pc+3]

	} else if typ == instructionTypeInput || typ == instructionTypeOutput {
		actualInstr.op1Mode = op1Mode(instr)
		if typ == instructionTypeInput && actualInstr.op1Mode == opModeImmediate {
			panic(fmt.Sprintf("got opModeImmediate for instruction that writes to address of 3rd operand, instr: %d", instr))
		}
		actualInstr.op1 = c.memory[c.pc+1]

	} else if typ == instructionTypeJumpIfTrue || typ == instructionTypeJumpIfFalse {
		actualInstr.op1Mode = op1Mode(instr)
		actualInstr.op2Mode = op2Mode(instr)

		actualInstr.op1 = c.memory[c.pc+1]
		actualInstr.op2 = c.memory[c.pc+2]

	} else if typ == instructionTypeAdjustRelativeBase {
		actualInstr.op1Mode = op1Mode(instr)
		actualInstr.op1 = c.memory[c.pc+1]

	} else if typ == instructionTypeHalt {

	} else {
		panic(fmt.Sprintf("encountered unknown instruction type %d", typ))
	}

	return actualInstr
}

func op1Mode(instr int) opMode {
	return opMode((instr / 100) % 10)
}

func op2Mode(instr int) opMode {
	return opMode((instr / 1000) % 10)
}

func op3Mode(instr int) opMode {
	return opMode((instr / 10000) % 10)
}

// returnrs whether the computer can continue running or if it has reached the halt instruction
func (c *computer) run() {
	c.running = true

	for c.running {

		fmt.Fprintf(computerStateWriter, "pc=%d rb=%d\n", c.pc, c.relativeBase)

		instrInt := c.memory[c.pc]

		fmt.Fprintln(instructionWriter, instrInt)

		instr := c.parseInstructionInt(instrInt)

		fmt.Fprintln(instructionWriter, instr)

		c.executeInstruction(instr)
	}

	fmt.Fprintln(computerStateWriter, "computer stopped")
}

// returns whether we needed input to execute an instruction but there was no input ready
func (c *computer) executeInstruction(instr instruction) {
	switch instr.typ {
	case instructionTypeAdd:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		var writeAddr int
		if instr.op3Mode == opModePosition {
			writeAddr = instr.op3
		} else if instr.op3Mode == opModeRelative {
			writeAddr = c.relativeBase + instr.op3
		}

		c.setValueAtAddress(writeAddr, v1+v2)

		c.pc += 4

	case instructionTypeMultiply:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		var writeAddr int
		if instr.op3Mode == opModePosition {
			writeAddr = instr.op3
		} else if instr.op3Mode == opModeRelative {
			writeAddr = c.relativeBase + instr.op3
		}

		c.setValueAtAddress(writeAddr, v1*v2)

		c.pc += 4

	case instructionTypeInput:
		input := c.getInput()

		switch instr.op1Mode {
		case opModePosition:
			c.setValueAtAddress(instr.op1, input)
		case opModeRelative:
			c.setValueAtAddress(c.relativeBase+instr.op1, input)
		default:
			panic(fmt.Sprintf("invalid op1 mode %d for input instruction", instr.op1Mode))
		}

		c.pc += 2

	case instructionTypeOutput:
		output := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)

		c.produceOutput(output)

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

		var writeAddr int
		if instr.op3Mode == opModePosition {
			writeAddr = instr.op3
		} else if instr.op3Mode == opModeRelative {
			writeAddr = c.relativeBase + instr.op3
		}

		if v1 < v2 {
			c.setValueAtAddress(writeAddr, 1)
		} else {
			c.setValueAtAddress(writeAddr, 0)
		}

		c.pc += 4

	case instructionTypeEquals:
		v1 := c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)
		v2 := c.loadValueBasedOnOpMode(instr.op2, instr.op2Mode)

		var writeAddr int
		if instr.op3Mode == opModePosition {
			writeAddr = instr.op3
		} else if instr.op3Mode == opModeRelative {
			writeAddr = c.relativeBase + instr.op3
		}

		if v1 == v2 {
			c.setValueAtAddress(writeAddr, 1)
		} else {
			c.setValueAtAddress(writeAddr, 0)
		}

		c.pc += 4

	case instructionTypeAdjustRelativeBase:
		c.relativeBase += c.loadValueBasedOnOpMode(instr.op1, instr.op1Mode)

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
	if a >= len(c.memory) {
		c.growMemory(a + 1) // we must grow to a+1 because that would make a the last valid index
	}
	v := c.memory[a]
	fmt.Fprintf(memoryAccessWriter, "read [%d] = %d\n", a, v)
	return v
}

// this function sets the value located in the computer's memory at address a to value v
// this is necessary because our computer must have memory beyond the initial program
// the memory of the computer is grown everytime we try to access a value that is currently out of bounds
// memory that has not been accessed yet is zeroed
func (c *computer) setValueAtAddress(a, v int) {
	if a >= len(c.memory) {
		c.growMemory(a + 1) // we must grow to a+1 because that would make a the last valid index
	}
	c.memory[a] = v
	fmt.Fprintf(memoryAccessWriter, "write [%d] = %d\n", a, v)
}

// grows the computer's memory to size n, so that the last valid index is n-1
func (c *computer) growMemory(n int) {
	newMemory := make([]int, n)
	copy(newMemory, c.memory)
	c.memory = newMemory
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %s\n", time.Since(start))
	}(time.Now())

	initialMemory := readInitialMemoryState()

	fmt.Println("Part 1:")
	doPart1(initialMemory)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(initialMemory)
}

func doPart1(initialMemory []int) {
	initialMemoryCopy := make([]int, len(initialMemory))
	copy(initialMemoryCopy, initialMemory)

	c := computer{
		memory: initialMemoryCopy,
		getInput: func() int {
			return 1
		},
		produceOutput: func(v int) {
			fmt.Printf("output: %d\n", v)
		},
	}

	c.run()
}

func doPart2(initialMemory []int) {
	initialMemoryCopy := make([]int, len(initialMemory))
	copy(initialMemoryCopy, initialMemory)

	c := computer{
		memory: initialMemoryCopy,
		getInput: func() int {
			return 2
		},
		produceOutput: func(v int) {
			fmt.Printf("output: %d\n", v)
		},
	}

	c.run()
}

func readInitialMemoryState() []int {
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
