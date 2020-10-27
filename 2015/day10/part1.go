package main

import "fmt"
import "strconv"

func main() {
	str := "1321131112"

	doChallenge(str)

	fmt.Println(seeAndSayNStepsLen(str, 40))
	
}

func doChallenge(str string) {
	for i := 0; i < 40; i++ {
		str = seeAndSay(str)
	}

	fmt.Println(len(str))

	for i := 0; i < 10; i++ {
		str = seeAndSay(str)
	}

	fmt.Println(len(str))
}

func seeAndSayNStepsLen(s string, n int) int {
	buf1 := make([]byte, 0, len(s) * 2 * n)
	buf2 := []byte(s)
	
	curBuf, otherBuf := buf1, buf2
	
	
	for i := 0; i < n; i++ {
		var cur byte
		var ct int64
		curBuf = curBuf[:1]
		for strIdx := 0; strIdx < len(otherBuf); strIdx++ {
			if cur == 0 {
				cur = otherBuf[strIdx]
				ct = 1
			} else if otherBuf[strIdx] != cur {
				curBuf = strconv.AppendInt(curBuf, ct, 10)
				curBuf = append(curBuf, cur)

				cur = otherBuf[strIdx]
				ct = 1
			} else {
				ct++
			}	
		}
		
		curBuf = strconv.AppendInt(curBuf, ct, 10)
		curBuf = append(curBuf, cur)
		
		curBuf, otherBuf = otherBuf, curBuf
	}
	
	return len(otherBuf)
}

func seeAndSay(s string) string {
	buf := make([]byte, 0, len(s)*2)

	var cur, ct byte
	for i := 0; i < len(s); i++ {
		if cur == 0 {
			cur = s[i]
			ct = 1
		} else if s[i] != cur {
			buf = strconv.AppendInt(buf, int64(ct), 10)
			buf = append(buf, cur)

			cur = s[i]
			ct = 1
		} else {
			ct++
		}
	}

	buf = strconv.AppendInt(buf, int64(ct), 10)
	buf = append(buf, cur)

	return string(buf)
}
