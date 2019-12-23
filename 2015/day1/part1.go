package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("provide input file")
	}

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	bufR := bufio.NewReader(f)

	fileBytes, err := ioutil.ReadAll(bufR)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	f.Close()

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
}
