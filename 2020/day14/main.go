package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	instrs := readInstrs()

	fmt.Println("Part 1:")
	doPart1(instrs)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(instrs)
}

func doPart1(instrs []instruction) {
	memory := make(map[int]int, len(instrs))

	var mask string

	for _, instr := range instrs {
		if instr.typ == instructionTypeSetMask {
			mask = instr.mask
		} else if instr.typ == instructionTypeSetMem {

			addr := instr.addr
			val := instr.val
			val = applyMask1(val, mask)

			memory[addr] = val

		} else {
			panic(fmt.Sprintf("invalid instruction type %d\n", instr.typ))
		}
	}

	var sum int
	for _, v := range memory {
		sum += v
	}

	fmt.Printf("Sum of values in memory: %d\n", sum)
}

func applyMask1(val int, mask string) int {
	bit := 35
	for _, r := range mask {
		if r == '0' {
			val &= ^(1 << bit)
		} else if r == '1' {
			val |= (1 << bit)
		}

		bit--
	}

	return val
}

func doPart2(instrs []instruction) {
	memory := make(map[int]int, len(instrs))

	var mask string

	for _, instr := range instrs {
		if instr.typ == instructionTypeSetMask {
			mask = instr.mask
		} else if instr.typ == instructionTypeSetMem {

			addr := instr.addr
			addrsToWrite := applyMask2(addr, mask)
			val := instr.val

			for _, addr := range addrsToWrite {
				memory[addr] = val
			}

		} else {
			panic(fmt.Sprintf("invalid instruction type %d\n", instr.typ))
		}
	}

	var sum int
	for _, v := range memory {
		sum += v
	}

	fmt.Printf("Sum of values in memory: %d\n", sum)
}

func applyMask2(addr int, mask string) []int {
	return applyMask2Help(addr, mask, 35)
}

func applyMask2Help(addr int, mask string, bit int) []int {
	if bit == -1 {
		return []int{addr}
	}

	addrs := make([]int, 0, 8)

	for i, r := range mask {
		if r == '1' {
			addr |= (1 << bit)

		} else if r == 'X' {
			{
				newAddr := addr & ^(1 << bit)
				addrs = append(addrs, applyMask2Help(newAddr, mask[i+1:], bit-1)...)
			}
			{
				newAddr := addr | (1 << bit)
				addrs = append(addrs, applyMask2Help(newAddr, mask[i+1:], bit-1)...)
			}
			return addrs
		}

		bit--
	}

	// this is only reached if the last bit of mask is a 1 or a 0
	addrs = append(addrs, addr)

	return addrs
}

type instructionType int

const (
	instructionTypeSetMask instructionType = iota + 1
	instructionTypeSetMem
)

type instruction struct {
	typ instructionType

	mask string // if typ == instructionTypeSetMask

	addr, val int // if typ == instructionTypeSetMem
}

func readInstrs() []instruction {
	scanner := bufio.NewScanner(os.Stdin)

	instrs := make([]instruction, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()

		lineWords := strings.Split(line, " ")

		var instr instruction
		if line[:3] == "mem" {
			instr.typ = instructionTypeSetMem

			{
				memParts := strings.Split(lineWords[0], "[")
				addrAndBrace := memParts[1]
				instr.addr, _ = strconv.Atoi(addrAndBrace[:len(addrAndBrace)-1])
			}
			instr.val, _ = strconv.Atoi(lineWords[2])

		} else {
			instr.typ = instructionTypeSetMask

			instr.mask = lineWords[2]
		}

		instrs = append(instrs, instr)
	}

	return instrs
}
