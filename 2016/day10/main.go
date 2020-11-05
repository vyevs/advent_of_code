package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instructionType int

const (
	instructionTypeValueToBot instructionType = iota
	instructionTypeBotDecides
)

type targetType int

const (
	targetTypeBot targetType = iota
	targetTypeOutput
)

type instruction struct {
	typ instructionType

	// for instructionTypeValueToBot
	value     int
	targetBot int

	// for instructionTypeBotDecides
	sourceBot                   int
	lowTargetTyp, highTargetTyp targetType
	lowTarget, highTarget       int
}

func main() {
	instrs := readInInstructions()

	executeInstructions(instrs)
}

func readInInstructions() []instruction {
	scanner := bufio.NewScanner(os.Stdin)

	instrs := make([]instruction, 0, 512)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		var instr instruction
		if words[0] == "value" {
			instr = parseInstructionTypeValueToBot(words)
		} else {
			instr = parseInstructionTypeBotDecides(words)
		}

		instrs = append(instrs, instr)
	}

	return instrs
}

// value 5 goes to bot 2
// 0     1 2    3  4   5
func parseInstructionTypeValueToBot(instrWords []string) instruction {
	var instr instruction
	instr.typ = instructionTypeValueToBot
	instr.value, _ = strconv.Atoi(instrWords[1])
	instr.targetBot, _ = strconv.Atoi(instrWords[5])
	return instr
}

// bot 2 gives low to bot 1 and high to bot 0
// 0   1 2     3   4  5   6 7   8    9  10  11
func parseInstructionTypeBotDecides(instrWords []string) instruction {
	var instr instruction
	instr.typ = instructionTypeBotDecides

	instr.sourceBot, _ = strconv.Atoi(instrWords[1])

	if instrWords[5] == "bot" {
		instr.lowTargetTyp = targetTypeBot
	} else {
		instr.lowTargetTyp = targetTypeOutput
	}

	if instrWords[10] == "bot" {
		instr.highTargetTyp = targetTypeBot
	} else {
		instr.highTargetTyp = targetTypeOutput
	}

	instr.lowTarget, _ = strconv.Atoi(instrWords[6])
	instr.highTarget, _ = strconv.Atoi(instrWords[11])

	return instr
}

type bot struct {
	v1, v2 int
}

func executeInstructions(instrs []instruction) {

	outputs := make(map[int]int, 512) // maps the output to the value it was given, if any
	bots := make(map[int]bot, 512)    // maps a bot's number to the bot itself

	alreadyExecuted := make([]bool, len(instrs))

	executedInstruction := true
	for executedInstruction {
		executedInstruction = false

		for i, instr := range instrs {
			if alreadyExecuted[i] {
				continue
			}
			executed := executeInstruction(instr, bots, outputs)
			if executed {
				executedInstruction = true
				alreadyExecuted[i] = true
			}
		}
	}

	log.Printf("bot state: %v", bots)
	log.Printf("final output state: %v", outputs)
}

// returns whether the instruction was successfully executed
func executeInstruction(instr instruction, bots map[int]bot, outputs map[int]int) bool {
	if instr.typ == instructionTypeValueToBot {
		giveValueToBot(instr.value, instr.targetBot, bots)
		return true

	} else if instr.typ == instructionTypeBotDecides {
		return executeInstructionTypeBotDecides(instr, bots, outputs)

	} else {
		panic(fmt.Sprintf("unknown instruction type %d", instr.typ))
	}
}

func giveValueToBot(v, targetBot int, bots map[int]bot) {
	bot, haveBot := bots[targetBot]

	if haveBot {
		if bot.v1 == -1 {
			bot.v1 = v
		} else if bot.v2 == -1 {
			bot.v2 = v
		} else {
			panic(fmt.Sprintf("attempted to give a value to a bot that has 2 values, value: %d, target bot: %v", v, targetBot))
		}
	} else {
		bot.v1 = v
		bot.v2 = -1
	}

	bots[targetBot] = bot
}

func executeInstructionTypeBotDecides(instr instruction, bots map[int]bot, outputs map[int]int) bool {
	bot, haveBot := bots[instr.sourceBot]
	// we can't execute the instruction if the bot doesn't have any values
	if !haveBot {
		return false
	}

	// we can't execute the instruction if the bot only has 1 value
	if bot.v1 == -1 || bot.v2 == -1 {
		return false
	}

	lowPtr := ptrToIntMin(&bot.v1, &bot.v2)
	highPtr := ptrToIntMax(&bot.v1, &bot.v2)

	low := *lowPtr
	high := *highPtr

	if low == 17 && high == 61 {
		log.Printf("bot %d compared microchips 17 and 61", instr.sourceBot)
	}

	*lowPtr = -1
	*highPtr = -1

	if instr.lowTargetTyp == targetTypeBot {
		giveValueToBot(low, instr.lowTarget, bots)
	} else {
		_, alreadyHaveOutput := outputs[instr.lowTarget]
		if alreadyHaveOutput {
			panic(fmt.Sprintf("output %d already has a value", instr.lowTarget))
		}
		outputs[instr.lowTarget] = low
	}

	if instr.highTargetTyp == targetTypeBot {
		giveValueToBot(high, instr.highTarget, bots)
	} else {
		_, alreadyHaveOutput := outputs[instr.highTarget]
		if alreadyHaveOutput {
			panic(fmt.Sprintf("output %d already has a value", instr.highTarget))
		}
		outputs[instr.highTarget] = high
	}

	return true
}

func ptrToIntMin(a, b *int) *int {
	if *a < *b {
		return a
	}
	return b
}

func ptrToIntMax(a, b *int) *int {
	if *a > *b {
		return a
	}
	return b
}
