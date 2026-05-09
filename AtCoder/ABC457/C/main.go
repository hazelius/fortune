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

	n, k := readInt(), readInt()-1
	ass := make([][]int, n)
	for i := range ass {
		l := readInt()
		ass[i] = make([]int, l)
		for j := range ass[i] {
			ass[i][j] = readInt()
		}
	}
	cs := make([]int, n)
	for i := range cs {
		cs[i] = readInt()
	}
	for i := 0; i < n; i++ {
		lng := cs[i] * len(ass[i])
		if k < lng {
			k = k % len(ass[i])
			fmt.Fprint(out, ass[i][k])
			return
		}
		k -= lng
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
