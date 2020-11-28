package main

import (
	"bufio"
	"container/list"
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
	instrs          []int
	pc              int // program counter
	running         bool
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

// returnrs whether the computer can continue running or if it has reached the halt instruction
func (c *computer) run() {
	c.running = true
	c.waitingForInput = false

	log.Printf("computer %c is running", c.name)

	for c.running {
		instrInt := c.instrs[c.pc]

		instr := c.parseInstructionInt(instrInt)

		c.executeInstruction(instr)
	}
}

func (c *computer) reset() {
	c.pc = 0
}

// returns whether we needed input to execute an instruction but there was no input ready
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
		inputElem := c.inputQueue.Front()

		if inputElem == nil {
			log.Printf("computer %c stopped, waiting for input", c.name)
			c.running = false
			c.waitingForInput = true
		} else {
			v := c.inputQueue.Remove(inputElem).(int)
			log.Printf("computer %c got input %d", c.name, v)
			log.Printf("computer %c input queue after getting input: %s", c.name, listString(c.inputQueue))
			c.instrs[instr.op1] = v
			c.pc += 2
		}

	case instructionTypeOutput:
		var output int
		if instr.op1Mode == opModeImmediate {
			output = instr.op1
		} else if instr.op1Mode == opModePosition {
			output = c.instrs[instr.op1]
		} else {
			panic(fmt.Sprintf("unknown op mode %d", instr.op1Mode))
		}

		c.outputTarget.inputQueue.PushBack(output)
		log.Printf("computer %c produced value %d", c.name, output)

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
		log.Printf("computer %c reached halt instruction", c.name)

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

	instrs := readInputInstructions()

	//doPart1(instrs)
	doPart2(instrs)
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

func doPart1(instrs []int) {
	highestSignal := determineHighestSignalThatCanBeSentToThrusters(instrs)

	log.Printf("highest signal that can be sent to thrusters: %d", highestSignal)
}

func doPart2(instrs []int) {
	highestSignal := determineHighestSignalThatCanBeSentToThrustersLooping(instrs)

	log.Printf("highest signal that can be sent to thrusters after looping amplifiers: %d", highestSignal)
}

func determineHighestSignalThatCanBeSentToThrusters(instrs []int) int {
	allPossiblePhaseSettings := generateAllPossiblePermutations([5]int{0, 1, 2, 3, 4})

	var highestSignal int
	for _, phaseSettings := range allPossiblePhaseSettings {
		signal := determineThrusterSignalForPhaseSequence(instrs, phaseSettings)
		if signal > highestSignal {
			highestSignal = signal
		}
	}

	return highestSignal
}

// returns the signal that is sent to the thrusters after running program instrs with the phase sequence seq
func determineThrusterSignalForPhaseSequence(instrs []int, phaseSettings [5]int) int {
	c := computer{
		instrs: instrs,
	}
	c.outputTarget = &c

	// the input signal to the first amplifier is always 0
	c.inputQueue.PushBack(0)

	for _, phaseSetting := range phaseSettings {
		// order of inputs to the computer is phase setting THEN the input signal
		c.inputQueue.PushFront(phaseSetting)
		c.run()

		c.reset()
	}

	thrustersSignal := c.inputQueue.Front().Value.(int)

	return thrustersSignal
}

func determineHighestSignalThatCanBeSentToThrustersLooping(instrs []int) int {
	allPossiblePhaseSettings := generateAllPossiblePermutations([5]int{5, 6, 7, 8, 9})

	var highestSignal int
	for _, phaseSettings := range allPossiblePhaseSettings {
		signal := determineThrusterSignalForPhaseSequenceLooping(instrs, phaseSettings)
		if signal > highestSignal {
			highestSignal = signal
		}
	}

	return highestSignal
}

func determineThrusterSignalForPhaseSequenceLooping(instrs []int, phaseSettings [5]int) int {
	var computers [5]computer
	for i := range computers {
		c := &computers[i]
		c.inputQueue = list.New()
		c.name = rune('A' + i)
		instrsCopy := make([]int, len(instrs))
		copy(instrsCopy, instrs)
		c.instrs = instrsCopy
		c.outputTarget = &computers[(i+1)%5]

		c.inputQueue.PushBack(phaseSettings[i])
	}

	computers[0].inputQueue.PushBack(0)

	for i, c := range computers {
		log.Printf("computer %c is outputting to computer %c", c.name, computers[(i+1)%5].name)
	}
	for _, c := range computers {
		log.Printf("computer %c initial queue state: %s", c.name, listString(c.inputQueue))
	}

	for i := 0; ; i++ {
		cIdx := i % 5
		c := computers[cIdx]

		log.Printf("computer %c input queue state: %s", c.name, listString(c.inputQueue))
		time.Sleep(1 * time.Second)

		c.run()

		if !c.waitingForInput {
			log.Printf("computer %c finished", c.name)
			if cIdx == 4 {
				break
			}
		}
	}

	thrustersSignal := computers[0].inputQueue.Front().Value.(int)

	return thrustersSignal
}

func generateAllPossiblePermutations(ints [5]int) [][5]int {

	permutations := make([][5]int, 0, 1024)

	var recurse func()

	var used [5]bool
	var nextIdx int
	var permutation [5]int

	recurse = func() {
		if nextIdx == 5 {
			permutations = append(permutations, permutation)
			return
		}

		for i, v := range ints {
			if !used[i] {
				used[i] = true
				permutation[nextIdx] = v

				nextIdx++
				recurse()
				nextIdx--

				used[i] = false
			}
		}
	}

	recurse()

	return permutations
}

func listString(l *list.List) string {
	var buf strings.Builder
	for e := l.Front(); e != nil; e = e.Next() {
		buf.WriteString(strconv.Itoa(e.Value.(int)))
		buf.WriteRune(' ')
	}
	return buf.String()
}
