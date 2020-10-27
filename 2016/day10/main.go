package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	instructions := parseInput(f)
	f.Close()

	for _, instruction := range instructions {
		fmt.Printf("%+v\n", instruction)
	}
}

func processInstructions(instructions []instruction) map[int]int {
	processed := make([]bool, len(instructions))

	type bot struct {
		lowChip  int
		highChip int
	}

	bots := make(map[int]bot, len(instructions)*2)
	outputs := make(map[int]int, len(instructions)*2)

	for {
		var processedAtLeastOne bool

		for i, instruction := range instructions {
			if !processed[i] {

				if instruction.typ == toBot {
					chipVal := instruction.lowTarget
					bot, ok := bots[instruction.highTarget]
					if ok {
						if chipVal > bot.lowChip {
							bot.highChip = chipVal
						} else {
							bot.lowChip, bot.highChip = chipVal, bot.lowChip
						}
					} else {
						bot.lowChip = chipVal
					}

					bots[instruction.highTarget] = bot
				} else { // instruction.typ == fromBot

					if instruction.lowTargetTyp == targetBot {

					}
				}

				processed[i] = true
				processedAtLeastOne = true
			}
		}

		if !processedAtLeastOne {
			break
		}
	}

	return outputs
}

type instructionType int

const (
	toBot instructionType = iota
	fromBot
)

type targetType int

const (
	targetBot targetType = iota
	targetOutput
)

type instruction struct {
	typ instructionType

	fromBot int

	lowTargetTyp  targetType
	highTargetTyp targetType

	lowTarget  int // if typ == toBot this is the chip value
	highTarget int // if typ == toBot this is the target bot
}

func parseInput(r io.Reader) []instruction {
	instructions := make([]instruction, 0, 512)

	scanner := bufio.NewScanner(r)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var instruction instruction
		var err error

		if tokens[0] == "value" {
			instruction.typ = toBot

			{
				chipValStr := tokens[1]

				instruction.lowTarget, err = strconv.Atoi(chipValStr)
				if err != nil {
					log.Fatalf("invalid chip value %q on line %d", chipValStr, lineNum)
				}
			}

			{
				targetStr := tokens[5]

				instruction.highTarget, err = strconv.Atoi(targetStr)
				if err != nil {
					log.Fatalf("invalid bot number %q on line %d", targetStr, lineNum)
				}
			}

		} else { // tokens[0] == "bot"
			instruction.typ = fromBot

			{
				fromBotStr := tokens[1]

				instruction.fromBot, err = strconv.Atoi(fromBotStr)
				if err != nil {
					log.Fatalf("invalid bot number %q on line %d", fromBotStr, lineNum)
				}
			}

			{
				lowTargetTypStr := tokens[5]

				if lowTargetTypStr == "bot" {
					instruction.lowTargetTyp = targetBot
				} else {
					instruction.lowTargetTyp = targetOutput
				}
			}

			{
				lowTargetStr := tokens[6]

				instruction.lowTarget, err = strconv.Atoi(lowTargetStr)
				if err != nil {
					log.Fatalf("invalid low target number %q on line %d", lowTargetStr, lineNum)
				}
			}

			{
				highTargetTypStr := tokens[10]

				if highTargetTypStr == "bot" {
					instruction.highTargetTyp = targetBot
				} else {
					instruction.highTargetTyp = targetOutput
				}
			}

			{
				highTargetStr := tokens[6]

				instruction.highTarget, err = strconv.Atoi(highTargetStr)
				if err != nil {
					log.Fatalf("invalid low target number %q on line %d", highTargetStr, lineNum)
				}
			}
		}

		instructions = append(instructions, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return instructions
}
