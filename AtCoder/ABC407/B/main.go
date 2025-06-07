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

	x, y := readInt(), readInt()
	cnt := 0
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i+j >= x || abs(i-j) >= y {
				cnt++
			}
		}
	}

	fmt.Fprint(out, float64(cnt)/36.0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	run(os.Stdin, os.Stdout)
}
