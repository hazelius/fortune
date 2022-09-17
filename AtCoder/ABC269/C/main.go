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

	n := readInt()
	bits := make([]int, 0)
	cnt := 0
	for i := n; i > 0; i = i >> 1 {
		cnt++
		if i%2 > 0 {
			bits = append(bits, cnt)
		}
	}

	m := 1
	for i := 0; i < len(bits); i++ {
		m *= 2
	}
	for i := 0; i < m; i++ {
		ans := 0
		for j, v := range bits {
			if (i>>j)%2 > 0 {
				ans += pow(2, v-1)
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func pow(a, b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r *= a
	}
	return r
}

func main() {
	run(os.Stdin, os.Stdout)
}
