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
	hs := make([]int, n)
	for i := range hs {
		hs[i] = readInt()
	}

	t := 0
	for _, h := range hs {
		for t%3 > 0 {
			t++
			if t%3 == 0 {
				h -= 3
			} else {
				h--
			}
			if h <= 0 {
				break
			}
		}

		t += (h / 5) * 3
		switch h % 5 {
		case 1:
			t++
		case 2:
			t += 2
		case 3, 4:
			t += 3
		}
	}

	fmt.Fprint(out, t)
}

func main() {
	run(os.Stdin, os.Stdout)
}
