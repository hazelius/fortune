package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	s := readString()
	ans := 1

	l := len(s)
	for i := 1; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			sa := j - i - 1
			ji := i - sa
			if ji < 0 || s[j] != s[ji] {
				break
			}

			t := 2 * (sa + 1)
			if ans < t {
				ans = t
			}
		}

		for j := i + 1; j < l; j++ {
			sa := j - i
			ji := i - sa
			if ji < 0 || s[j] != s[ji] {
				break
			}

			t := 2*sa + 1
			if ans < t {
				ans = t
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
