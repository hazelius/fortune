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

	s := "wbwbwwbwbwbw"
	wc, bc := 0, 0
	for _, c := range s {
		if c == 'w' {
			wc++
		} else {
			bc++
		}
	}

	w, b := readInt(), readInt()
	if w == 0 && b == 0 {
		fmt.Fprint(out, "Yes")
		return
	}

	cnt := min(w/wc, b/bc)
	w -= wc * cnt
	b -= bc * cnt

	s = s + s
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			tw, tb := 0, 0
			for k := i; k < j; k++ {
				if s[k] == 'w' {
					tw++
				} else {
					tb++
				}
			}
			if tw == w && tb == b {
				fmt.Fprint(out, "Yes")
				return
			}
		}
	}

	fmt.Fprint(out, "No")
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
