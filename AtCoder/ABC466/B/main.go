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
	ms := make([]int, m)
	for i := range ms {
		ms[i] = -1
	}

	for i := 0; i < n; i++ {
		c, s := readInt()-1, readInt()
		if ms[c] < s {
			ms[c] = s
		}
	}

	ans := fmt.Sprintf("%v", ms)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
