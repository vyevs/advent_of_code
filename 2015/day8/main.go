package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vyevs/vtools"
)

func main() {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	var sum int
	var encSum int

	for _, line := range lines {

		codeLen := codeLength(line)

		sum += len(line) - codeLen

		encLen := encodedLen(line)

		encSum += encLen - len(line)
	}

	fmt.Println(sum)
	fmt.Println(encSum)
}

func codeLength(s string) int {
	var ct int
	for i := 1; i < len(s)-1; {
		ct++
		if s[i] == '\\' {
			if s[i+1] == 'x' {
				i += 4
			} else {
				i += 2
			}
		} else {
			i++
		}
	}
	return ct
}

func encodedLen(s string) int {
	len := 2
	for _, r := range s {
		if r == '"' || r == '\\' {
			len += 2
		} else {
			len += 1
		}
	}

	return len
}

func isHexDigit(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f')
}
