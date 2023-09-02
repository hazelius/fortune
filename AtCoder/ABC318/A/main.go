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

	n, m, p := readInt(), readInt(), readInt()
	ans := 0
	if n >= m {
		n -= m
		ans++
	} else {
		fmt.Fprint(out, ans)
		return
	}

	ans += n / p
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
