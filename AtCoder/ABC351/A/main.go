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

	ans := 0
	for i := 0; i < 9; i++ {
		ans += readInt()
	}
	for i := 0; i < 8; i++ {
		ans -= readInt()
	}
	fmt.Fprint(out, ans+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
