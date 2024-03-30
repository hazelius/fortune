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

	a, b, c := readInt(), readInt(), readInt()

	x, y := 0, 0
	for i := 0; i <= 60; i++ {
		if c&(1<<i) > 0 {
			if a == 0 && b == 0 {
				fmt.Fprint(out, -1)
				return
			}

			if a >= b {
				x |= 1 << i
				a--
			} else {
				y |= 1 << i
				b--
			}
		}
	}
	if a != b {
		fmt.Fprint(out, -1)
		return
	}

	for i := 0; i < 60; i++ {
		if a == 0 {
			break
		}
		if x&(1<<i) == 0 && y&(1<<i) == 0 {
			x |= 1 << i
			y |= 1 << i
			a--
		}
	}

	if a > 0 {
		fmt.Fprint(out, -1)
		return
	}

	fmt.Fprint(out, x, y)
	// fmt.Printf("%b %b %v\n", x, y, x^y)

}

func main() {
	run(os.Stdin, os.Stdout)
}
