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
		log.Fatal("provide an input file")
	}

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("open: %v", err)
	}

	bufR := bufio.NewReader(f)

	fileBytes, err := ioutil.ReadAll(bufR)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}

	type coord struct {
		x, y int
	}

	visited := make(map[coord]struct{}, 10000)

	var loc1 coord
	var loc2 coord
	visited[loc1] = struct{}{}

	for i, b := range fileBytes {

		var moved *coord
		if i%2 == 0 {
			moved = &loc1
		} else {
			moved = &loc2
		}

		switch b {
		case '>':
			moved.x++
		case '<':
			moved.x--
		case '^':
			moved.y++
		case 'v':
			moved.y--
		}

		visited[*moved] = struct{}{}
	}

	fmt.Println(len(visited))
}
