package main

import (
	"fmt"
	"os"

	"github.com/vyevs/vtools"
)

func main() {
	boxIDs, err := vtools.ReadLines(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: the checksum of the box IDs is %d\n", calculateChecksum(boxIDs))

	fmt.Printf("part 2: the common letters between the two correct box IDs are %q\n", getCommonLettersOfCorrectBoxIDs(boxIDs))
}

func calculateChecksum(boxIDs []string) int {
	var ct2s, ct3s int
	for _, it := range boxIDs {
		charCts := vtools.Counter(vtools.StrBytes(it))

		var seen2s, seen3s bool
		for _, ct := range charCts {
			if ct == 2 && !seen2s {
				ct2s++
				seen2s = true
			} else if ct == 3 && !seen3s {
				ct3s++
				seen3s = true
			}
		}
	}

	return ct2s * ct3s
}

func getCommonLettersOfCorrectBoxIDs(boxIDs []string) string {
	var cid1, cid2 string
	for i, id1 := range boxIDs {
		for _, id2 := range boxIDs[i+1:] {
			if equalButOneLetter(id1, id2) {
				cid1, cid2 = id1, id2
			}
		}
	}

	for i := range cid1 {
		c1, c2 := cid1[i], cid2[i]
		if c1 != c2 {
			return cid1[:i] + cid1[i+1:]
		}
	}

	panic("no such box IDs")
}

func equalButOneLetter(s1, s2 string) bool {
	var diffCt int
	for i := range s1 {
		c1, c2 := s1[i], s2[i]
		if c1 != c2 {
			diffCt++
			if diffCt == 2 {
				return false
			}

		}
	}

	return diffCt == 1
}
