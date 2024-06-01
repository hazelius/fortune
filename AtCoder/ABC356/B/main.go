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
	as := make([]int, m)
	for i := range as {
		as[i] = readInt()
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := readInt()
			as[j] -= x
		}
	}

	for _, v := range as {
		if v > 0 {
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
