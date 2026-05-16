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

	h, w := readInt(), readInt()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cnt := 4
			if i == 0 {
				cnt--
			}
			if i == h-1 {
				cnt--
			}
			if j == 0 {
				cnt--
			}
			if j == w-1 {
				cnt--
			}
			fmt.Fprint(out, cnt)
			if j < w-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
