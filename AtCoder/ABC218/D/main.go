package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type vector struct {
	a int
	b int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ps := make([][]int, n)
	for i := range ps {
		ps[i] = []int{readInt(), readInt()}
	}

	ym := make(map[vector]int)

	ans := 0
	for i, p := range ps {
		for j := 0; j < i; j++ {
			c := ps[j]
			if p[0] == c[0] {
				v := vector{p[1], c[1]}
				if p[1] > c[1] {
					v = vector{c[1], p[1]}
				}
				cnt, ok := ym[v]
				if ok {
					ans += cnt
				}
				ym[v]++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
