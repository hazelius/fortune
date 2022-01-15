package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var used map[int]bool

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	a, n := readInt(), readInt()
	used = make(map[int]bool)
	return f(a, 0, []int{n})
}

func f(a, cnt int, q []int) int {
	if len(q) == 0 {
		return -1
	}
	nextq := make([]int, 0)
	cnt++

	for _, n := range q {
		if n%a == 0 {
			next := n / a
			if next == 1 {
				return cnt
			}
			if !used[next] {
				used[next] = true
				nextq = append(nextq, next)
			}
		}
		if n < 10 {
			continue
		}

		strn := strconv.Itoa(n)
		strn = strn[1:] + strn[:1]
		if strn[:1] == "0" {
			continue
		}

		np, _ := strconv.Atoi(strn)
		if np == 1 {
			return cnt
		}
		if !used[np] {
			used[np] = true
			nextq = append(nextq, np)
		}
	}

	return f(a, cnt, nextq)
}

func main() {
	fmt.Println(run(os.Stdin))
}
