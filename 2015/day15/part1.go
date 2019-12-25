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

	cookies := parseInput(f)
	f.Close()

	const (
		totalCookies = 100
	)

	cookieCounts := make([]int, 0, len(cookies))
	var best int

	var fn func()
	fn = func() {
		if len(cookieCounts) == len(cookies)-1 {
			nLastCookie := totalCookies - sum(cookieCounts)
			cookieCounts = append(cookieCounts, nLastCookie)

			if totalCalories(cookies, cookieCounts) != 500 {
				cookieCounts = cookieCounts[:len(cookieCounts)-1]
				return
			}

			score := score(cookies, cookieCounts)
			if score > best {
				best = score
			}

			cookieCounts = cookieCounts[:len(cookieCounts)-1]

			return
		}

		for i := 0; i < totalCookies; i++ {
			cookieCounts = append(cookieCounts, i)
			fn()
			cookieCounts = cookieCounts[:len(cookieCounts)-1]
		}
	}

	fn()

	fmt.Println(best)

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

func parseInput(r io.Reader) []cookie {
	scanner := bufio.NewScanner(r)

	cookies := make([]cookie, 0, 64)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var cookie cookie

		cookie.name = tokens[0][:len(tokens[0])-1]
		cookie.capacity, _ = strconv.Atoi(tokens[2][:len(tokens[2])-1])
		cookie.durability, _ = strconv.Atoi(tokens[4][:len(tokens[4])-1])
		cookie.flavor, _ = strconv.Atoi(tokens[6][:len(tokens[6])-1])
		cookie.texture, _ = strconv.Atoi(tokens[8][:len(tokens[8])-1])
		cookie.calories, _ = strconv.Atoi(tokens[10])

		cookies = append(cookies, cookie)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return cookies
}

func sum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}
