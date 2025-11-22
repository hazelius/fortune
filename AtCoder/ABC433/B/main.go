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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	for i, a := range as {
		ans := -1
		for j := i - 1; j >= 0; j-- {
			if a < as[j] {
				ans = j + 1
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
