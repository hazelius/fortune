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

	k, q := -1, -1
	r, b, n := make([]int, 0), make([]int, 0), make([]int, 0)
	for i, v := range s {
		switch v {
		case 'K':
			k = i
			if len(r) != 1 {
				fmt.Fprint(out, "No")
				return
			}
		case 'Q':
			q = i
		case 'R':
			r = append(r, i)
		case 'B':
			b = append(b, i)
		case 'N':
			n = append(n, i)
		}
	}

	if k == -1 || q == -1 || len(r) < 2 || len(b) < 2 || len(n) < 2 {
		fmt.Fprint(out, "No")
		return
	}
	if b[0]%2 == b[1]%2 {
		fmt.Fprint(out, "No")
		return
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
