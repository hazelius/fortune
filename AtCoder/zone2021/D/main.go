package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s string) string {
	var ans []byte
	flg := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if string(c) == "R" {
			flg = !flg
			continue
		}
		if len(ans) == 0 {
			ans = []byte{c}
			continue
		}

		if flg {
			ans = append([]byte{c}, ans...)
		} else {
			ans = append(ans, c)
		}
	}

	l := len(ans)
	ans2 := make([]byte, l)
	ai := 0
	for i := 0; i < l; i++ {
		c := ans[i]
		if flg {
			c = ans[l-i-1]
		}

		if ai > 0 && c == ans2[ai-1] {
			ai--
			continue
		}
		ans2[ai] = c
		ai++
	}
	return string(ans2[:ai])
}

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
