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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	k := 1
	r, c := 0, n/2
	ans[r][c] = k
	for i := 0; i < n*n-1; i++ {
		k++
		r2 := (r + n - 1) % n
		c2 := (c + 1) % n
		if r2 < n && r2 >= 0 && c2 < n && c2 >= 0 && ans[r2][c2] == 0 {
			ans[r2][c2] = k
		} else {
			r2 = (r + 1) % n
			c2 = c
			ans[r2][c2] = k
		}
		r = r2
		c = c2
	}
	for _, v := range ans {
		ans := fmt.Sprintf("%v", v)
		fmt.Fprintln(out, ans[1:len(ans)-1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
