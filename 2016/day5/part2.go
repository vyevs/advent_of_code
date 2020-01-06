package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	prefix := []byte("cxdnnyjw")

	pw := make([]rune, 8)
	placed := make([]bool, 8)

	hasher := md5.New()

	md5Sum := make([]byte, 0, hex.EncodedLen(md5.Size))

	var charsPlaced int
	for i := 0; charsPlaced < 8; i++ {
		in := strconv.AppendInt(prefix, int64(i), 10)

		_, _ = hasher.Write(in)

		md5Sum := hasher.Sum(md5Sum)

		if md5Sum[0] == 0 && md5Sum[1] == 0 && (md5Sum[2]&0xF0) == 0 {
			position := rune(md5Sum[2] & 0xF)

			if position < 8 && !placed[position] {
				char := rune(md5Sum[3]&0xF0) >> 4

				if char >= 0 && char <= 9 {
					char += '0'
				} else {
					char -= 10
					char += 'a'
				}
				fmt.Printf("%q\n", char)
				pw[position] = char
				placed[position] = true

				charsPlaced++
			}
		}

		hasher.Reset()
	}

	fmt.Println(string(pw))
}
