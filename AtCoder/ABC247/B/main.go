package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	sts := make([][]string, n)
	stmap := make(map[string]int)
	for i := range sts {
		s, t := readString(), readString()
		sts[i] = []string{s, t}
		stmap[s]++
		stmap[t]++
	}

	for _, st := range sts {
		flg := false
		for _, v := range st {
			cnt := stmap[v]
			if cnt > 2 {
				continue
			}
			if cnt == 2 {
				if st[0] != st[1] {
					continue
				}
			}
			flg = true
		}
		if !flg {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
