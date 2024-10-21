package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vyevs/vtools"
)

func main() {
	dims := getInput()

	var totalSurfaceArea int
	var totalRibbon int

	for _, dim := range dims {
		l, w, h := dim[0], dim[1], dim[2]

		s1 := l * h
		s2 := l * w
		s3 := h * w

		totalSurfaceArea += 2*s1 + 2*s2 + 2*s3 + min(s1, min(s2, s3))
		totalRibbon += smallestPerimeterOfOneFace(l, h, w) + l*h*w
	}

	fmt.Println(totalSurfaceArea)
	fmt.Println(totalRibbon)
}

func getInput() [][3]int {
	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic("provide an input file")
	}

	out := make([][3]int, 0, 1024)

	for _, line := range lines {
		var l, w, h int
		n, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			log.Fatalf("Sscanf: %v", err)
		} else if n != 3 {
			log.Fatalf("line doesn't have 3 values")
		}

		out = append(out, [3]int{l, w, h})
	}

	return out
}

func smallestPerimeterOfOneFace(l, h, w int) int {
	return min(2*l+2*h, min(2*h+2*w, 2*l+2*w))
}
