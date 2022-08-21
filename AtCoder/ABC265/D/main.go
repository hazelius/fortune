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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, p, q, r := readInt(), readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	lidx, ridx, sum := 0, 0, 0
	t := p + q + r
	for {
		if sum == t {
			nidx := check(lidx, p, as)
			if nidx > 0 {
				nidx = check(nidx, q, as)
			}
			if nidx > 0 {
				fmt.Fprint(out, "Yes")
				return
			}
		}
		if sum <= t {
			if ridx >= len(as) {
				break
			}
			sum += as[ridx]
			ridx++
		} else {
			sum -= as[lidx]
			lidx++
		}
	}

	fmt.Fprint(out, "No")
}

func check(idx, t int, as []int) int {
	sum := 0
	for i := idx; i < len(as); i++ {
		sum += as[i]
		if sum == t {
			return i + 1
		}
		if sum > t {
			return -1
		}
	}

	return -1
}

func main() {
	run(os.Stdin, os.Stdout)
}
