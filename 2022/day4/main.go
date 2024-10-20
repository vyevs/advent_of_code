package main

import (
	"fmt"
	"io"
)

func main() {
	in := getInput()
	part1(in)
	part2(in)
}

func part1(vs [][4]int) {
	var ct int
	for _, v := range vs {
		if oneRangeCoversTheOther(v[0], v[1], v[2], v[3]) {
			ct += 1
		}
	}
	
	fmt.Printf("%d of the pairs have one range capturing the other\n", ct)
}


func oneRangeCoversTheOther(v1, v2, v3, v4 int) bool {
	return (v1 <= v3 && v2 >= v4) || (v3 <= v1 && v4 >= v2)
}


func part2(vs [][4]int) {
	var ct int
	for _, v := range vs {
		if rangesOverlap(v[0], v[1], v[2], v[3]) {
			ct += 1
		}
	}
	
	fmt.Printf("%d of the ranges overlap\n", ct)
}





func rangesOverlap(v1, v2, v3, v4 int) bool {
	return (v2 >= v3 && v1 <= v3) || // end of [v1, v2] overlaps beginning of [v3, v4]
	(v1 <= v4 && v2 >= v4) || // end of [v3, v4] overlaps beginning of [v1, v2]
	oneRangeCoversTheOther(v1, v2, v3, v4) // one of the ranges is entirely inside the other
}


func getInput() [][4]int {
	out := make([][4]int, 0, 1024)
	for {
		var v1, v2, v3, v4 int
		_, err := fmt.Scanf("%d-%d,%d-%d\n", &v1, &v2, &v3, &v4)
		if err == io.EOF {
			break
		}
		
		line := [4]int{v1, v2, v3, v4}
		out = append(out, line)
	}
	
	return out
}