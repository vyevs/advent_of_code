package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	triangles := parseInput(f)
	f.Close()

	fmt.Printf("%d total triangles\n", len(triangles))

	part1Answer := numPossibleTriangles(triangles)
	fmt.Printf("part1 answer: %d\n", part1Answer)

	vertTris := verticalizeTriangles(triangles)
	part2Answer := numPossibleTriangles(vertTris)
	fmt.Printf("part1 answer: %d\n", part2Answer)
}

func numPossibleTriangles(tris []triangle) int {
	var possible int
	for _, tri := range tris {
		sortedSidesTri := sortTriangleSides(tri)
		if sortedSidesTri.s1+sortedSidesTri.s2 > sortedSidesTri.s3 {
			possible++
		}
	}

	return possible
}

func verticalizeTriangles(tris []triangle) []triangle {
	result := make([]triangle, 0, len(tris))

	var row, col int

	for col < 3 {
		if row >= len(tris) {
			row = 0
			col++
		} else {

			var t triangle

			if col == 0 {
				t.s1 = tris[row].s1
				t.s2 = tris[row+1].s1
				t.s3 = tris[row+2].s1
			} else if col == 1 {
				t.s1 = tris[row].s2
				t.s2 = tris[row+1].s2
				t.s3 = tris[row+2].s2
			} else if col == 2 {
				t.s1 = tris[row].s3
				t.s2 = tris[row+1].s3
				t.s3 = tris[row+2].s3
			}

			result = append(result, t)

			row += 3
		}
	}

	return result
}

func sortTriangleSides(t triangle) triangle {
	if t.s1 > t.s3 {
		t.s1, t.s3 = t.s3, t.s1
	}
	if t.s2 > t.s3 {
		t.s2, t.s3 = t.s3, t.s2
	}

	return t
}

type triangle struct {
	s1, s2, s3 int
}

func parseInput(r io.Reader) []triangle {
	scanner := bufio.NewScanner(r)

	triangles := make([]triangle, 0, 4096)

	for scanner.Scan() {
		line := scanner.Text()

		lenStrs := strings.Split(line, " ")

		for i := 0; i < len(lenStrs); {
			if lenStrs[i] == "" {
				lenStrs[i], lenStrs[len(lenStrs)-1] = lenStrs[len(lenStrs)-1], lenStrs[i]
				lenStrs = lenStrs[:len(lenStrs)-1]
			} else {
				i++
			}
		}

		var tri triangle

		tri.s1, _ = strconv.Atoi(lenStrs[0])
		tri.s2, _ = strconv.Atoi(lenStrs[1])
		tri.s3, _ = strconv.Atoi(lenStrs[2])

		triangles = append(triangles, tri)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return triangles
}
