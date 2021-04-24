package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, s string, tabs [][]int) string {
	flips := []byte(s)

	flg := false
	for _, tab := range tabs {
		switch tab[0] {
		case 1:
			a := tab[1] - 1
			b := tab[2] - 1
			if flg {
				if a >= n {
					a -= n
				} else {
					a += n
				}
				if b >= n {
					b -= n
				} else {
					b += n
				}
			}
			flips[a], flips[b] = flips[b], flips[a]
		case 2:
			flg = !flg
		}
	}

	if !flg {
		return string(flips)
	}

	return string(flips[n:]) + string(flips[:n])
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	s := readString()
	q := readInt()
	tabs := make([][]int, q)
	for i := range tabs {
		tabs[i] = []int{readInt(), readInt(), readInt()}
	}
	fmt.Println(run(n, s, tabs))
}
