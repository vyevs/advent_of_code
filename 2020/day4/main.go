package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	passports := readPassports()

	fmt.Println("Part 1:")
	doPart1(passports)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(passports)
}

func doPart1(passports []passport) {
	var nValid int
	for _, pp := range passports {
		if pp.isValid1() {
			nValid++
		}
	}

	fmt.Printf("%d passports are valid\n", nValid)
}

func (pp passport) isValid1() bool {
	return pp.byr != "" && pp.iyr != "" && pp.eyr != "" && pp.hgt != "" && pp.hcl != "" &&
		pp.ecl != "" && pp.pid != ""
}

func doPart2(passports []passport) {
	var nValid int
	for _, pp := range passports {
		if pp.isValid2() {
			nValid++
		}
	}

	fmt.Printf("%d passports are valid\n", nValid)
}

func (pp passport) isValid2() bool {
	// byr
	{
		if len(pp.byr) != 4 || !isAllDigits(pp.byr) {
			return false
		}

		byrInt, _ := strconv.Atoi(pp.byr)
		if byrInt < 1920 || byrInt > 2002 {
			return false
		}
	}

	// iyr
	{
		if len(pp.iyr) != 4 || !isAllDigits(pp.iyr) {
			return false
		}
		iyrInt, _ := strconv.Atoi(pp.iyr)
		if iyrInt < 2010 || iyrInt > 2020 {
			return false
		}
	}

	// eyr
	{
		if len(pp.eyr) != 4 || !isAllDigits(pp.eyr) {
			return false
		}
		eyrInt, _ := strconv.Atoi(pp.eyr)
		if eyrInt < 2020 || eyrInt > 2030 {
			return false
		}
	}

	// hgt
	{
		hl := len(pp.hgt)

		if hl < 4 || !isAllDigits(pp.hgt[:hl-2]) {
			return false
		}

		hgtInt, _ := strconv.Atoi(pp.hgt[:hl-2])

		if pp.hgt[hl-2:] == "cm" {
			if hgtInt < 150 || hgtInt > 193 {
				return false
			}

		} else if pp.hgt[hl-2:] == "in" {
			if hgtInt < 59 || hgtInt > 76 {
				return false
			}

		} else {
			return false
		}
	}

	// hcl
	{
		if len(pp.hcl) != 7 || pp.hcl[0] != '#' {
			return false
		}
		for i := 1; i < len(pp.hcl); i++ {
			c := pp.hcl[i]
			if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
				return false
			}
		}
	}

	// ecl
	{
		switch pp.ecl {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		default:
			return false
		}
	}

	// pid
	{
		if len(pp.pid) != 9 || !isAllDigits(pp.pid) {
			return false
		}
	}

	return true
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func readPassports() []passport {
	scanner := bufio.NewScanner(os.Stdin)

	passports := make([]passport, 0, 128)
	var pp passport
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			passports = append(passports, pp)
			pp = passport{}
		} else {
			kvs := strings.Split(line, " ")

			for _, kv := range kvs {
				split := strings.Split(kv, ":")

				k := split[0]
				v := split[1]

				switch k {
				case "byr":
					pp.byr = v
				case "iyr":
					pp.iyr = v
				case "eyr":
					pp.eyr = v
				case "hgt":
					pp.hgt = v
				case "hcl":
					pp.hcl = v
				case "ecl":
					pp.ecl = v
				case "pid":
					pp.pid = v
				case "cid":
					pp.cid = v
				}
			}
		}
	}

	passports = append(passports, pp)

	return passports
}
