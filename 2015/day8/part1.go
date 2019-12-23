package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	scanner := bufio.NewScanner(f)

	var sum int
	var encSum int

	for scanner.Scan() {
		line := scanner.Text()

		codeLen := codeLength(line)

		sum += len(line) - codeLen

		encLen := encodedLen(line)

		encSum += encLen - len(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
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
