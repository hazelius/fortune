package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type ar struct {
	v   int
	cnt int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	q := readInt()
	as := make([]ar, q)
	idx := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			c, x := readInt(), readInt()
			as[i] = ar{x, c}
		case 2:
			ans := 0
			k := readInt()
			for k > 0 {
				if as[idx].cnt <= k {
					ans += as[idx].cnt * as[idx].v
					k -= as[idx].cnt
					as[idx].cnt = 0
					idx++
				} else {
					ans += k * as[idx].v
					as[idx].cnt -= k
					k = 0
				}
			}
			fmt.Fprintln(out, ans)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
