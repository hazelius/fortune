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

	n, k := readInt(), readInt()
	ans := make([]int, n)
	f(0, 0, n, k, ans, out)
}

func f(idx, s, n, k int, arg []int, out io.Writer) {
	if idx == n-1 {
		if (k-s)%n == 0 {
			arg[idx] = (k - s) / n
			strans := fmt.Sprintf("%v", arg)
			fmt.Fprintln(out, strans[1:len(strans)-1])
		}
		return
	}

	for cnt := 0; s+cnt*(idx+1) <= k; cnt++ {
		arg[idx] = cnt
		f(idx+1, s+cnt*(idx+1), n, k, arg, out)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
