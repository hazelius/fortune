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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := 0
	l, r := 1, 2
	for r <= n {
		fmt.Fprintln(out, "?", l, r)
		res := readString()
		if res == "Yes" {
			r++
		} else {
			ans += r - l - 1
			l++
			if r == l {
				r++
			}
		}
	}
	for ; l < n; l++ {
		ans += r - l - 1
	}

	fmt.Fprintln(out, "!", ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
