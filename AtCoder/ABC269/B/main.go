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

	a, b, c, d := 0, 0, 0, 0
	for i := 0; i < 10; i++ {
		s := readString()
		flg := false
		for j, l := range s {
			if l == '#' {
				flg = true
				if c == 0 {
					c = j + 1
				}
			} else {
				if flg && d == 0 {
					d = j
				}
			}
		}
		if flg {
			if a == 0 {
				a = i + 1
			}
		} else {
			if a > 0 && b == 0 {
				b = i
			}
		}
	}
	if b == 0 {
		b = 10
	}
	if d == 0 {
		d = 10
	}
	fmt.Fprintf(out, "%v %v\n%v %v", a, b, c, d)
}

func main() {
	run(os.Stdin, os.Stdout)
}
