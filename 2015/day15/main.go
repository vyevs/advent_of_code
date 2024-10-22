package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/vyevs/vtools"
)

const (
	totalCookies = 100
)

func main() {
	defer vtools.TimeIt(time.Now(), "everything")

	lines, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	cookies := vtools.Map(lines, parseLine)

	fmt.Printf("part 1: %d\n", doIt(cookies, false))
	fmt.Printf("part 2: %d\n", doIt(cookies, true))
}

func doIt(cookies []cookie, part2 bool) int {
	cookieCounts := vtools.NewStack[int](len(cookies))
	var best int

	cookiesLeft := totalCookies

	var fn func()
	fn = func() {
		// If we've picked the amounts of all but one cookie...
		if len(cookieCounts) == len(cookies)-1 {
			cookieCounts.Push(cookiesLeft)
			defer cookieCounts.Pop()

			if part2 && totalCalories(cookies, cookieCounts) != 500 {
				return
			}

			score := score(cookies, cookieCounts)
			if score > best {
				best = score
			}

			return
		}

		if cookiesLeft == 0 {
			cookieCounts.Push(0)
			fn()
			_ = cookieCounts.Pop()

		} else {
			for i := range cookiesLeft + 1 {
				cookieCounts.Push(i)
				cookiesLeft -= i
				fn()
				_ = cookieCounts.Pop()
				cookiesLeft += i
			}
		}
	}

	fn()

	return best
}

func totalCalories(cookies []cookie, counts []int) int {
	var calories int
	for i, cookie := range cookies {
		calories += cookie.calories * counts[i]
	}
	return calories
}

func score(cookies []cookie, counts []int) int {
	var capSum int
	for i, cookie := range cookies {
		capSum += cookie.capacity * counts[i]
	}
	if capSum <= 0 {
		return 0
	}

	var durSum int
	for i, cookie := range cookies {
		durSum += cookie.durability * counts[i]
	}
	if durSum <= 0 {
		return 0
	}

	var flavorSum int
	for i, cookie := range cookies {
		flavorSum += cookie.flavor * counts[i]
	}
	if flavorSum <= 0 {
		return 0
	}

	var texSum int
	for i, cookie := range cookies {
		texSum += cookie.texture * counts[i]
	}
	if texSum <= 0 {
		return 0
	}

	return capSum * durSum * flavorSum * texSum
}

type cookie struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var re = regexp.MustCompile(`^(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)

func parseLine(line string) cookie {
	groups := re.FindStringSubmatch(line)

	c := cookie{
		name:       groups[1],
		capacity:   atoi(groups[2]),
		durability: atoi(groups[3]),
		flavor:     atoi(groups[4]),
		texture:    atoi(groups[5]),
		calories:   atoi(groups[6]),
	}
	return c
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
