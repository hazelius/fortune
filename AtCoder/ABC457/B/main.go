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
	ass := make([][]int, n)
	for i := range ass {
		l := readInt()
		ass[i] = make([]int, l)
		for j := range ass[i] {
			ass[i][j] = readInt()
		}
	}
	x, y := readInt(), readInt()
	fmt.Fprint(out, ass[x-1][y-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
