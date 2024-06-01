package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
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
	sc.Split(bufio.ScanWords)

	n, m, k := readInt(), readInt(), readInt()
	as := make([]uint, m)
	rs := make([]bool, m)
	for i := 0; i < m; i++ {
		c := readInt()
		a := uint(0)
		for j := 0; j < c; j++ {
			a |= 1 << (readInt() - 1)
		}
		as[i] = a
		rs[i] = readString() == "o"
	}

	ans := 0
	for i := uint(0); i < (1 << n); i++ {
		flg := true
		for j, a := range as {
			bc := bits.OnesCount(i & a)
			b := rs[j]
			if bc < k && b {
				flg = false
				break
			}
			if bc >= k && !b {
				flg = false
				break
			}
		}
		if flg {
			ans++
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
