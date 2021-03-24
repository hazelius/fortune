package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

type k struct {
	v int
	b bool
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s, t string) int {
	minimam := len(s)
	tlen := len(t)
	for i := 0; i <= len(s)-tlen; i++ {
		m := 0
		st := s[i : i+tlen]
		// fmt.Println(st)
		for j := 0; j < tlen; j++ {
			if m >= minimam {
				break
			}
			if st[j:j+1] != t[j:j+1] {
				m++
			}
		}
		minimam = m
	}
	return minimam
}

func main() {
	sc.Split(bufio.ScanWords)
	s, t := readString(), readString()
	fmt.Println(run(s, t))
}
