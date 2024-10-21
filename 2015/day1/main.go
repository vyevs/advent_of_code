package main

import (
	"fmt"
	"os"
)

func main() {
	if err := doIt(); err != nil {
		fmt.Printf("error: %v", err)
	}
}

func doIt() error {
	fileBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	var floor int
	var firstEnteredBasement int
	for i, b := range fileBytes {
		if b == '(' {
			floor++
		} else {
			floor--
			if floor == -1 && firstEnteredBasement == 0 {
				firstEnteredBasement = i + 1
			}
		}
	}

	fmt.Println(floor)
	fmt.Println(firstEnteredBasement)

	return nil
}
