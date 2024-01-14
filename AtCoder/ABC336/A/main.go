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
	ans := "L"
	for i := 0; i < n; i++ {
		ans += "o"
	}
	fmt.Fprintf(out, "%vng", ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
