package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	prefix := "iwrupvqb"

	hasher := md5.New()

	cipher := make([]byte, 0, 16)
	hexEncoding := make([]byte, 32)

	for i := 0; ; i++ {

		str := fmt.Sprintf("%s%d", prefix, i)

		_, _ = hasher.Write([]byte(str))

		cipher := hasher.Sum(cipher)

		_ = hex.Encode(hexEncoding, cipher)

		if startsWith6Zeroes(hexEncoding) {
			fmt.Printf("%d\n%s\n%s\n", i, str, hexEncoding)
			os.Exit(0)
		}

		hasher.Reset()
	}
}

func startsWith6Zeroes(bs []byte) bool {
	for i := 0; i < 6; i++ {
		if bs[i] != '0' {
			return false
		}
	}

	return true
}
