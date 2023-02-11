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

	n, m := readInt(), readInt()
	cs := make([]int, m)
	for i := range cs {
		c := readInt()
		for j := 0; j < c; j++ {
			a := readInt()
			cs[i] += 1 << (a - 1)
		}
	}
	l := 1<<n - 1

	ans := 0
	for i := 1; i < 1<<m; i++ {
		tmp := 0
		for j := 0; j < m; j++ {
			if i&(1<<j) > 0 {
				tmp |= cs[j]
			}
		}
		if tmp == l {
			ans++
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
