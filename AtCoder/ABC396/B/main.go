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
	as := make([]int, 0)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			as = append(as, readInt())
		case 2:
			x := 0
			if len(as) > 0 {
				x = as[len(as)-1]
				as = as[:len(as)-1]
			}
			fmt.Fprintln(out, x)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
