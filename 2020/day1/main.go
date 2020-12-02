package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ns := readInNumbers()

	fmt.Println("Part 1:")
	doPart1(ns)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(ns)
}

func doPart1(ns []int) {
	nsLen := len(ns)
	n1, n2 := -1, -1
	for i := 0; i < nsLen; i++ {
		for j := 0; j < nsLen; j++ {
			if i != j && ns[i]+ns[j] == 2020 {
				n1 = ns[i]
				n2 = ns[j]
				break
			}
		}
	}

	if n1 == -1 {
		panic("there are no 2 numbers that add to 2020")
	}

	fmt.Printf("n1 = %d n2 = %d\n", n1, n2)
	fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
}

func doPart2(ns []int) {
	nsLen := len(ns)
	n1, n2, n3 := -1, -1, -1
	for i := 0; i < nsLen; i++ {
		for j := 0; j < nsLen; j++ {
			for k := 0; k < nsLen; k++ {
				if allIntsDiffer(i, j, k) && ns[i]+ns[j]+ns[k] == 2020 {
					n1 = ns[i]
					n2 = ns[j]
					n3 = ns[k]
					break
				}
			}
		}
	}

	if n1 == -1 {
		panic("there are no 3 numbers that add to 2020")
	}

	fmt.Printf("n1=%d n2=%d n3=%d\n", n1, n2, n3)
	fmt.Printf("%d * %d * %d = %d\n", n1, n2, n3, n1*n2*n3)
}

func allIntsDiffer(i1, i2, i3 int) bool {
	return i1 != i2 && i1 != i3 && i2 != i3
}

func readInNumbers() []int {
	scanner := bufio.NewScanner(os.Stdin)

	ns := make([]int, 0, 1024)
	for lineNum := 1; scanner.Scan(); lineNum++ {
		line := scanner.Text()

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("input line %d is invalid %q", lineNum, line))
		}

		ns = append(ns, n)
	}

	return ns
}
