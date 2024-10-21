package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileBytes, err := os.ReadFile(os.Args[1])
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
