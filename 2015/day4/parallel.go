package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	maxProcs := runtime.GOMAXPROCS(0)

	fmt.Printf("GOMAXPROCS: %d\n", maxProcs)

	prefix := "iwrupvqb"

	startTime := time.Now()

	for i := 0; i < maxProcs; i++ {
		go func(start int) {

			hasher := md5.New()

			cipher := make([]byte, 0, 16)

			strBuf := make([]byte, 8, 100)

			for i := range prefix {
				strBuf[i] = prefix[i]
			}

			for i := start; ; i += maxProcs {
				strBuf := strconv.AppendInt(strBuf, int64(i), 10)

				_, _ = hasher.Write(strBuf)

				cipher := hasher.Sum(cipher)

				if cipher[0] == 0 && cipher[1] == 0 && cipher[2] == 0 && cipher[3] == 0 {
					encoding := hex.EncodeToString(cipher)
					fmt.Printf("%d\n%s\n%s\n", i, strBuf, encoding)
					fmt.Printf("took %s\n", time.Since(startTime))
					os.Exit(0)
				}

				hasher.Reset()
			}

		}(i)
	}

	select {}
}
