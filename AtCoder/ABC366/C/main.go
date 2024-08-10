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

	q := readInt()
	xmap := make(map[int]int)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			xmap[x]++
		case 2:
			x := readInt()
			cnt := xmap[x]
			if cnt == 1 {
				delete(xmap, x)
			} else {
				xmap[x] = cnt - 1
			}
		case 3:
			fmt.Fprintln(out, len(xmap))
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
