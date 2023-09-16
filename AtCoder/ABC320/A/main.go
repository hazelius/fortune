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

	a, b := readInt(), readInt()

	ans := multi(a, b) + multi(b, a)

	fmt.Fprint(out, ans)
}

func multi(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
