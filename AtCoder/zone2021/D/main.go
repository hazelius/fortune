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
			if c == ans[0] {
				ans = ans[1:]
			} else {
				ans = append([]byte{c}, ans...)
			}
		} else {
			if c == ans[len(ans)-1] {
				ans = ans[:len(ans)-1]
			} else {
				ans = append(ans, c)
			}
		}
	}

	if !flg {
		return string(ans)
	}

	l := len(ans)
	ans2 := make([]byte, l)
	for i := 0; i < l; i++ {
		ans2[i] = ans[l-i-1]
	}
	return string(ans2)
}

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
