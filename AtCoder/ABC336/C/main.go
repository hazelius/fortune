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
	n--
	ans := 0
	for i := 1; ; i *= 10 {
		t := n % 5
		ans += i * t
		if n < 5 {
			break
		}
		n /= 5
	}

	fmt.Fprint(out, ans*2)
}

func main() {
	run(os.Stdin, os.Stdout)
}
