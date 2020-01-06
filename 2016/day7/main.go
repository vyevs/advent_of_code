package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])

	ips := parseInput(f)
	f.Close()

	var tlsCount, sslCount int
	for _, ip := range ips {
		if supportsTLS(ip) {
			tlsCount++
		}
		if supportsSSL(ip) {
			sslCount++
		}
	}

	fmt.Printf("ips that support TLS: %d\n", tlsCount)
	fmt.Printf("ips that support SSL: %d\n", sslCount)
}

func supportsSSL(ip ip) bool {
	exteriorABAs := allABAs(ip.exteriorSeqs)
	hypernetBABs := allABAs(ip.hypernetSeqs)

	for _, aba := range exteriorABAs {
		for _, bab := range hypernetBABs {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}
	return false
}

func allABAs(strs []string) []string {
	abas := make([]string, 0, 128)
	for _, s := range strs {
		for i := 0; i < len(s)-2; i++ {
			if (s[i] == s[i+2]) && (s[i] != s[i+1]) {
				abas = append(abas, s[i:i+2]) // only use 2 chars to represent the aba as "ab"
			}
		}
	}
	return abas
}

func supportsTLS(ip ip) bool {
	return hasABBA(ip.exteriorSeqs) && !hasABBA(ip.hypernetSeqs)
}

func hasABBA(strs []string) bool {
	for _, s := range strs {
		for i := 0; i < len(s)-3; i++ {
			if (s[i] != s[i+1]) && (s[i] == s[i+3]) && (s[i+1] == s[i+2]) {
				return true
			}
		}
	}
	return false
}

type ip struct {
	exteriorSeqs []string
	hypernetSeqs []string
}

func parseInput(r io.Reader) []ip {
	scanner := bufio.NewScanner(r)

	ips := make([]ip, 0, 2048)

	for scanner.Scan() {
		line := scanner.Text()

		ip := ip{
			exteriorSeqs: make([]string, 0, 8),
			hypernetSeqs: make([]string, 0, 8),
		}

		var inBracket bool
		for len(line) > 0 {
			if inBracket {
				closedBracketIndex := strings.Index(line, "]")

				ip.hypernetSeqs = append(ip.hypernetSeqs, line[:closedBracketIndex])

				line = line[closedBracketIndex+1:]
				inBracket = false

			} else {
				openBracketIndex := strings.Index(line, "[")

				if openBracketIndex == -1 {
					ip.exteriorSeqs = append(ip.exteriorSeqs, line)
					line = line[len(line):]
				} else {
					ip.exteriorSeqs = append(ip.exteriorSeqs, line[:openBracketIndex])
					line = line[openBracketIndex+1:]
				}

				inBracket = true
			}
		}

		ips = append(ips, ip)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return ips
}
