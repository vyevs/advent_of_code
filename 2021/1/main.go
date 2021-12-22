package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	measurements := getMeasurements()
	{
		nIncreases := numIncreases(measurements)
		fmt.Printf("1: There are %d increases\n", nIncreases)
	}
	{
		nIncreases := nIncreasesSlidingWindow(measurements)
		fmt.Printf("2: There are %d increases\n", nIncreases)
	}
}

func numIncreases(ns []int) int {
	var nIncreases int
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[i-1] {
			nIncreases++
		}
	}
	return nIncreases
}

func nIncreasesSlidingWindow(ns []int) int {
	var nIncreases int
	winSum := ns[0] + ns[1] + ns[2]
	for i := 3; i < len(ns); i++ {
		oldSum := winSum
		winSum -= ns[i-3]
		winSum += ns[i]
		if winSum > oldSum {
			nIncreases++
		}
	}
	return nIncreases
}


func getMeasurements() []int {
	scanner := bufio.NewScanner(os.Stdin)
	
	
	nums := make([]int, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		
		num, _ := strconv.Atoi(line)
		
		nums = append(nums, num)
	}
	
	return nums
}