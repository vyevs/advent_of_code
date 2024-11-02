package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vyevs/vtools"
)

func main() {
	exampleTimes := []uint64{7, 15, 30}
	exampleDistances := []uint64{9, 40, 200}

	fmt.Printf("part 1 example: %d\n", productOfNumberOfWaysToWin(exampleTimes, exampleDistances))

	times := []uint64{49, 78, 79, 80}
	distances := []uint64{298, 1185, 1066, 1181}
	fmt.Printf("part 1: %d\n", productOfNumberOfWaysToWin(times, distances))

	{
		exampleTime := joinu64s(exampleTimes)
		exampleDistance := joinu64s(exampleDistances)
		fmt.Printf("part 2 example: %d\n", nWaysToWin(exampleTime, exampleDistance))
	}

	time := joinu64s(times)
	distance := joinu64s(distances)
	fmt.Printf("part 2: %d\n", nWaysToWin(time, distance))
}

func productOfNumberOfWaysToWin(times, distances []uint64) uint64 {
	p := uint64(1)
	for i, t := range times {
		p *= nWaysToWin(t, distances[i])
	}
	return p
}

func nWaysToWin(time, distance uint64) uint64 {
	var n uint64
	for holdTime := uint64(1); holdTime < time; holdTime++ {
		speed := holdTime
		distTravelled := speed * (time - holdTime)
		if distTravelled > distance {
			n++
		}
	}
	return n
}

func joinu64s(vs []uint64) uint64 {
	intStrs := vtools.MapSlice(vs, func(v uint64) string {
		return strconv.FormatUint(v, 10)
	})
	str := strings.Join(intStrs, "")

	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return v
}
