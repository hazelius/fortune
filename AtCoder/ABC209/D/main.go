package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	abm := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := readInt(), readInt()
		if ab, ok := abm[a]; ok {
			ab = append(ab, b)
			abm[a] = ab
		} else {
			abm[a] = []int{b}
		}
		if ab, ok := abm[b]; ok {
			ab = append(ab, a)
			abm[b] = ab
		} else {
			abm[b] = []int{a}
		}
	}

	dism := make(map[int]int)
	f(1, 0, abm, dism)

	ans := make([]string, q)
	for i := 0; i < q; i++ {
		c, d := readInt(), readInt()
		cdis, ddis := dism[c], dism[d]
		if cdis%2 == ddis%2 {
			ans[i] = "Town"
		} else {
			ans[i] = "Road"
		}
	}

	return strings.Join(ans, "\n")
}

func f(n, d int, abm map[int][]int, dism map[int]int) {
	for _, b := range abm[n] {
		if _, ok := dism[b]; ok {
			continue
		}
		dism[b] = d + 1
		f(b, d+1, abm, dism)
	}
}

func main() {
	fmt.Println(run(os.Stdin))
}
