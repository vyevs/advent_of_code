package main

import "fmt"

func main() {
	src := "zzy"
	pw := pw{
		bs:    make([]byte, 64),
		bsLen: 64,
		pwLen: len(src),
	}
	for i := range src {
		pw.bs[64-len(src)+i] = src[i]
	}

	fmt.Println(string(pw.bs[pw.bsLen-pw.pwLen:]))
}

type pw struct {
	bs    []byte
	bsLen int
	pwLen int
}

func next(pw pw) pw {

	var carry byte
	for i := pw.bsLen - 1; ; i-- {
		pw.bs[i]++
		if pw.bs[i] > 'z' {
			pw.bs[i] = 'a'
			carry = 1
		} else {
			break
		}
	}

	return pw
}
