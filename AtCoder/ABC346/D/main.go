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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()
	cs := make([]int, n)
	for i := range cs {
		cs[i] = readInt()
	}

	secost := make([]int, n)
	recost := make([]int, n)
	for i, c := range s {
		if c == '0' {
			if i%2 > 0 {
				recost[i] = cs[i]
			} else {
				secost[i] = cs[i]
			}
		} else {
			if i%2 > 0 {
				secost[i] = cs[i]
			} else {
				recost[i] = cs[i]
			}
		}

		if i > 0 {
			secost[i] += secost[i-1]
			recost[i] += recost[i-1]
		}
	}

	ans := -1
	for i := 0; i < n-1; i++ {
		cost1 := secost[i] + recost[n-1] - recost[i]
		if ans < 0 || cost1 < ans {
			ans = cost1
		}
		cost2 := recost[i] + secost[n-1] - secost[i]
		if ans < 0 || cost2 < ans {
			ans = cost2
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
