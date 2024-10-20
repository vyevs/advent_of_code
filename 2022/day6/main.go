package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	in := getInput()
	part1(in)
	part2(in)
}

func part1(in string) {
	idx := findPacketMarker(in)
	
	fmt.Printf("the packet marker ends after %d characters\n", idx)
}

func part2(in string) {
	idx := findMsgMarker(in)
	
	fmt.Printf("the message marker ends after %d characters\n", idx)
}

func findPacketMarker(in string) int {
	const packetMarkerLen = 4
	msgMrkrIdx := indexUniqRunSlice(in, packetMarkerLen)
	if msgMrkrIdx == -1 {
		panic(fmt.Sprintf("%s does not contain a message marker", in))
	}
	
	return msgMrkrIdx + packetMarkerLen
}

func findMsgMarker(in string) int {
	const msgMarkerLen = 14
	msgMrkrIdx := indexUniqRunSlice(in, msgMarkerLen)
	if msgMrkrIdx == -1 {
		panic(fmt.Sprintf("%s does not contain a message marker", in))
	}
	
	return msgMrkrIdx + msgMarkerLen
}

func indexUniqRunMap(in string, l int) int {
	end := len(in) - (l - 1)
	for i := 0; i < end; i++ {
		substr := in[i:i+l]
		
		uniqs := make(map[rune]struct{}, l)
		
		for _, r := range substr {
			uniqs[r] = struct{}{}
		}
		
		if len(uniqs) == l {
			return i
		}
	}
	
	return -1
}

func indexUniqRunSlice(in string, l int) int {
	charCts := make([]int, 26)
	var nUniq int
	for i := 0; i < l; i++ {
		c := in[i] - 'a'
		addChar(charCts, c, &nUniq)
	}
	if nUniq == l {
		return 0
	}
	
	end := len(in) - l + 1
	for i := 1; i < end; i++ {
		{
			c := in[i-1] - 'a'
			removeChar(charCts, c, &nUniq)
		}
		
		{
			c := in[i+l-1] - 'a'
			addChar(charCts, c, &nUniq)
		}
			
		if nUniq == l {
			return i
		}
	}
	
	return -1
}

func addChar(cts []int, c byte, nUniq *int) {
	cts[c]++
	if cts[c] == 1 {
		*nUniq++
	} else if cts[c] == 2 {
		*nUniq--
	}
}

func removeChar(cts []int, c byte, nUniq *int) {
	cts[c]--
	if cts[c] == 0 {
		*nUniq--
	} else if cts[c] == 1 {
		*nUniq++
	}
}

func getInput() string {
	in, _ := io.ReadAll(os.Stdin)
	return string(in)
}