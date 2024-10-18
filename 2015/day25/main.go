package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("that took %s\n", time.Since(start))
	}(time.Now())

	v := uint64(20151125)
	r, c := 1, 1
	for {
		if r == 3010 && c == 3019 {
			fmt.Printf("[%d %d] %d\n", r, c, v)
			break
		}

		v *= 252533
		v %= 33554393

		if r == 1 {
			r = c + 1
			c = 1
		} else {
			r--
			c++
		}
	}

}
