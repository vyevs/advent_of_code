package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	prefix := []byte("cxdnnyjw")

	pw := make([]rune, 0, 8)

	hasher := md5.New()

	md5Sum := make([]byte, 0, hex.EncodedLen(md5.Size))

	for i := 0; len(pw) != 8; i++ {
		in := strconv.AppendInt(prefix, int64(i), 10)

		_, _ = hasher.Write(in)

		md5Sum := hasher.Sum(md5Sum)

		if md5Sum[0] == 0 && md5Sum[1] == 0 && (md5Sum[2]&0xF0) == 0 {
			sixthChar := rune(md5Sum[2] & 0xF)
			if sixthChar >= 0 && sixthChar <= 9 {
				sixthChar += '0'
			} else {
				sixthChar -= 10
				sixthChar += 'a'
			}
			pw = append(pw, sixthChar)
		}

		hasher.Reset()
	}

	fmt.Println(string(pw))
}
