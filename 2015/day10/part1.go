package main

import "fmt"

func main() {
	str := "1321131112"

	for i := 0; i < 40; i++ {
		str = seeAndSay(str)
	}

	fmt.Println(len(str))

	str = "1321131112"

	for i := 0; i < 50; i++ {
		str = seeAndSay(str)
	}

	fmt.Println(len(str))
}

func seeAndSay(s string) string {
	buf := make([]byte, 0, len(s)*2)

	var cur, ct byte
	for i := 0; i < len(s); i++ {
		if cur == 0 {
			cur = s[i]
			ct = 1
		} else if s[i] != cur {
			for ; ct > 0; ct /= 10 {
				buf = append(buf, ct%10+'0')
			}
			buf = append(buf, cur)

			cur = s[i]
			ct = 1
		} else {
			ct++
		}
	}

	for ; ct > 0; ct /= 10 {
		buf = append(buf, ct%10+'0')
	}
	buf = append(buf, cur)

	return string(buf)
}
