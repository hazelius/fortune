package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var used = []bool{}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	mapn := make(map[int][]int)
	for i := 0; i < q; i++ {
		l, r := readInt(), readInt()
		l--
		mapn[l] = append(mapn[l], r)
		mapn[r] = append(mapn[r], l)
	}
	used = make([]bool, n+1)
	if f(0, n, mapn) {
		return "Yes"
	}

	return "No"
}

func f(s, g int, mp map[int][]int) bool {
	used[s] = true
	ts, ok := mp[s]
	if !ok {
		return false
	}
	for _, t := range ts {
		if t == g {
			return true
		}
		if used[t] {
			continue
		}
		if f(t, g, mp) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(run(os.Stdin))
}
