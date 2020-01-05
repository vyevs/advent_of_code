package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	infos := parseInput(f)
	f.Close()

	realRooms := filterFakeRooms(infos)

	var sectorSum int
	for _, info := range realRooms {
		sectorSum += info.sector
	}

	fmt.Printf("sum of sectors = %d\n", sectorSum)

	for _, room := range realRooms {
		rotatedName := rotateName(room.name, room.sector)

		if rotatedName == "northpole object storage" {
			fmt.Printf("sector of northpole object room: %d\n", room.sector)
		}
	}
}

func rotateName(name string, amt int) string {
	buf := make([]rune, 0, len(name))

	amt %= 26

	for _, r := range name {
		if r == '-' {
			buf = append(buf, ' ')

		} else {
			rotatedRune := r + rune(amt)

			if rotatedRune > 'z' {
				rotatedRune = 'a' + (rotatedRune - 'z' - 1)
			}

			buf = append(buf, rotatedRune)
		}
	}

	return string(buf)
}

func filterFakeRooms(infos []roomInfo) []roomInfo {

	realRooms := make([]roomInfo, 0, len(infos))

	for _, info := range infos {
		if realRoom(info) {
			realRooms = append(realRooms, info)
		}
	}

	return realRooms
}

func realRoom(info roomInfo) bool {
	type charCounter struct {
		char  rune
		count int
	}

	var charCounts [26]charCounter
	for i := 0; i < len(charCounts); i++ {
		charCounts[i].char = rune('a' + i)
	}

	for _, c := range info.name {
		if c != '-' {
			charCounts[c-'a'].count++
		}
	}

	for i := 0; i < 5; i++ {
		for j := len(charCounts) - 1; j > i; j-- {
			curCount := &charCounts[j]
			prevCount := &charCounts[j-1]

			if curCount.count > prevCount.count ||
				(curCount.count == prevCount.count && curCount.char < prevCount.char) {

				*curCount, *prevCount = *prevCount, *curCount
			}
		}
	}

	for i := 0; i < 5; i++ {
		if charCounts[i].char != rune(info.checksum[i]) {
			return false
		}
	}

	return true
}

type roomInfo struct {
	name     string
	sector   int
	checksum string
}

func parseInput(r io.Reader) []roomInfo {
	scanner := bufio.NewScanner(r)

	infos := make([]roomInfo, 0, 256)

	for scanner.Scan() {
		line := scanner.Text()

		var info roomInfo

		splitOnBracket := strings.Split(line, "[")

		nameAndSectorID := splitOnBracket[0]

		lastDashIndex := strings.LastIndex(nameAndSectorID, "-")

		info.name = nameAndSectorID[:lastDashIndex]

		sectorStr := nameAndSectorID[lastDashIndex+1:]

		info.sector, _ = strconv.Atoi(sectorStr)

		checksum := splitOnBracket[1]
		checksum = checksum[:len(checksum)-1]

		info.checksum = checksum

		infos = append(infos, info)
	}

	return infos
}
