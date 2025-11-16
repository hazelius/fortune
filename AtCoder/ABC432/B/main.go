package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

	x := readInt()
	xs := make([]int, 0)
	for {
		xs = append(xs, x%10)
		x /= 10
		if x == 0 {
			break
		}
	}

	sort.Ints(xs)
	for i, v := range xs {
		if v > 0 {
			xs[0], xs[i] = xs[i], xs[0]
			break
		}
	}

	for _, v := range xs {
		fmt.Fprint(out, v)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
