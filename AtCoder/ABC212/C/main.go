package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, m)
	for i := range bs {
		bs[i] = readInt()
	}
	sort.Ints(as)
	sort.Ints(bs)

	ret := int(1e9 + 1)
	ai, bi := 0, 0
	for {
		a, b := as[ai], bs[bi]
		d := abs(a - b)
		if d == 0 {
			return 0
		}
		if ret > d {
			ret = d
		}
		if a > b {
			bi++
			if bi >= len(bs) {
				break
			}
		} else {
			ai++
			if ai >= len(as) {
				break
			}
		}
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
