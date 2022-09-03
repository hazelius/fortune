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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	tmp := 0
	sum := 0
	for i := 0; i < m; i++ {
		tmp += as[i]
		sum += as[i] * (i + 1)
	}

	ans := sum
	for i := 1; i <= n-m; i++ {
		sum -= tmp
		tmp = tmp - as[i-1] + as[i+m-1]
		sum += as[i+m-1] * m

		if ans < sum {
			ans = sum
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
