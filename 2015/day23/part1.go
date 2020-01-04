package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	typ instructionType

	register  uint8 // 0 for A, 1 for B
	jmpOffset int   // for jmp, jie, jio
}

type instructionType int

const (
	instructionTypeHlf instructionType = iota
	instructionTypeTpl
	instructionTypeInc
	instructionTypeJmp
	instructionTypeJie
	instructionTypeJio
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	instructions := parseInput(f)
	f.Close()

	registerA, registerB := executeInstructions(instructions)

	fmt.Printf("register A after instructions: %s\n", registerA.String())
	fmt.Printf("register B after instructions: %s\n", registerB.String())
}

func executeInstructions(instructions []instruction) (big.Int, big.Int) {
	registerA := *big.NewInt(1)
	var registerB big.Int

	instructionIndex := 0

	for instructionIndex < len(instructions) {
		instruction := instructions[instructionIndex]

		switch instruction.typ {
		case instructionTypeHlf:
			var targetRegister *big.Int
			if instruction.register == 0 {
				targetRegister = &registerA
			} else if instruction.register == 1 {
				targetRegister = &registerB
			}

			targetRegister.Rsh(targetRegister, 1)

			instructionIndex++

		case instructionTypeTpl:
			var targetRegister *big.Int
			if instruction.register == 0 {
				targetRegister = &registerA
			} else if instruction.register == 1 {
				targetRegister = &registerB
			}

			targetRegister.Mul(targetRegister, big.NewInt(3))

			instructionIndex++

		case instructionTypeInc:
			var targetRegister *big.Int
			if instruction.register == 0 {
				targetRegister = &registerA
			} else if instruction.register == 1 {
				targetRegister = &registerB
			}

			targetRegister.Add(targetRegister, big.NewInt(1))

			instructionIndex++

		case instructionTypeJmp:
			instructionIndex += instruction.jmpOffset

		case instructionTypeJie:
			var targetRegister *big.Int
			if instruction.register == 0 {
				targetRegister = &registerA
			} else if instruction.register == 1 {
				targetRegister = &registerB
			}

			if targetRegister.Bit(0) == 0 {
				instructionIndex += instruction.jmpOffset
			} else {
				instructionIndex++
			}

		case instructionTypeJio:
			var targetRegister *big.Int
			if instruction.register == 0 {
				targetRegister = &registerA
			} else if instruction.register == 1 {
				targetRegister = &registerB
			}

			if targetRegister.IsUint64() && targetRegister.Uint64() == 1 {
				instructionIndex += instruction.jmpOffset
			} else {
				instructionIndex++
			}

		default:
			log.Fatalf("unknown instruction type %d", instruction.typ)
		}
	}

	return registerA, registerB
}

func parseInput(r io.Reader) []instruction {
	scanner := bufio.NewScanner(r)

	instructions := make([]instruction, 0, 128)

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var instruction instruction

		switch tokens[0] {
		case "hlf":
			instruction.typ = instructionTypeHlf
			if tokens[1][0] == 'b' {
				instruction.register = 1
			}

		case "tpl":
			instruction.typ = instructionTypeTpl
			if tokens[1][0] == 'b' {
				instruction.register = 1
			}

		case "inc":
			instruction.typ = instructionTypeInc
			if tokens[1][0] == 'b' {
				instruction.register = 1
			}

		case "jmp":
			instruction.typ = instructionTypeJmp
			offset, err := strconv.Atoi(tokens[1])
			if err != nil {
				log.Fatalf("error parsing offset amount on line %d", lineNum)
			}
			instruction.jmpOffset = offset

		case "jie":
			instruction.typ = instructionTypeJie
			if tokens[1][0] == 'b' {
				instruction.register = 1
			}
			offset, err := strconv.Atoi(tokens[2])
			if err != nil {
				log.Fatalf("error parsing offset amount on line %d", lineNum)
			}
			instruction.jmpOffset = offset

		case "jio":
			instruction.typ = instructionTypeJio
			if tokens[1][0] == 'b' {
				instruction.register = 1
			}
			offset, err := strconv.Atoi(tokens[2])
			if err != nil {
				log.Fatalf("error parsing offset amount on line %d", lineNum)
			}
			instruction.jmpOffset = offset

		default:
			log.Fatalf("unknown instruction %s on line %d", tokens[0], lineNum)
		}

		instructions = append(instructions, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return instructions
}
