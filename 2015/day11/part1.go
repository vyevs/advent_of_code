package main

import "fmt"

func main() {
	src := "cqjxjnds"

	pw := pwFromString(src)

	for pw = next(pw); !legalPassword(pw); pw = next(pw) {
	}

	fmt.Println(pw)

	for pw = next(pw); !legalPassword(pw); pw = next(pw) {
	}

	fmt.Println(pw)
}

func legalPassword(pw pw) bool {
	return !containsIllegalLetters(pw) && containsRun(pw) && contains2Pairs(pw)
}

func containsRun(pw pw) bool {
	for i := pw.bsLen - pw.pwLen; i < pw.bsLen-2; i++ {
		if pw.bs[i+1] == pw.bs[i]+1 && pw.bs[i+2] == pw.bs[i+1]+1 {
			return true
		}
	}
	return false
}

func containsIllegalLetters(pw pw) bool {
	for _, b := range pw.bs[pw.bsLen-pw.pwLen:] {
		if b == 'i' || b == 'o' || b == 'l' {
			return true
		}
	}
	return false
}

func contains2Pairs(pw pw) bool {
	var firstPair byte
	for i := pw.bsLen - pw.pwLen; i < pw.bsLen-1; {
		if pw.bs[i] == pw.bs[i+1] {
			if firstPair == 0 {
				firstPair = pw.bs[i]
				i += 2
			} else if pw.bs[i] != firstPair {
				return true
			}
		} else {
			i++
		}
	}
	return false
}

type pw struct {
	bs    []byte
	bsLen int
	pwLen int
}

func pwFromString(s string) pw {
	pw := pw{
		bs:    make([]byte, 64),
		bsLen: 64,
		pwLen: len(s),
	}
	for i := range s {
		pw.bs[64-len(s)+i] = s[i]
	}
	return pw
}

func (pw pw) String() string {
	return string(pw.bs[pw.bsLen-pw.pwLen:])
}

func next(pw pw) pw {

	carry := true
	for i := pw.bsLen - 1; i >= pw.bsLen-pw.pwLen && carry; i-- {
		pw.bs[i]++
		if pw.bs[i] > 'z' {
			pw.bs[i] = 'a'
		} else {
			carry = false
		}
	}

	if carry {
		pw.pwLen++
		pw.bs[pw.bsLen-pw.pwLen] = 'a'
	}

	return pw
}
