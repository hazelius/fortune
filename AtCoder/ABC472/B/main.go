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
	as := make([]int, n)
	sum := 0
	for i := range as {
		as[i] = readInt()
		sum += as[i]
	}

	ans := sum
	sum2 := 0
	for _, v := range as {
		sum2 += v
		tmp := abs(sum - sum2*2)
		if tmp < ans {
			ans = tmp
		}
	}

	fmt.Fprint(out, ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
