package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(f)

	var totalSurfaceArea int
	var totalRibbon int

	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		var l, h, w int
		n, err := fmt.Sscanf(line, "%dx%dx%d", &l, &h, &w)
		if err != nil {
			log.Fatalf("Sscanf: %v", err)
		} else if n != 3 {
			log.Fatalf("line %d doesn't have 3 values", lineNum)
		}

		s1 := l * h
		s2 := l * w
		s3 := h * w

		totalSurfaceArea += 2*s1 + 2*s2 + 2*s3 + min(s1, min(s2, s3))
		totalRibbon += smallestPerimeterOfOneFace(l, h, w) + l*h*w
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	fmt.Println(totalSurfaceArea)
	fmt.Println(totalRibbon)
}

func smallestPerimeterOfOneFace(l, h, w int) int {
	return min(2*l+2*h, min(2*h+2*w, 2*l+2*w))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
