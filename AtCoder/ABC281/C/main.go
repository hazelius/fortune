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

	n, t := readInt(), readInt()
	as := make([]int, n)
	sum := 0
	for i := range as {
		as[i] = readInt()
		if sum >= 0 && sum < t {
			sum += as[i]
		} else {
			sum = -1
		}
	}

	if sum > 0 {
		t %= sum
	}

	for i, a := range as {
		if a < t {
			t -= a
			continue
		}
		fmt.Fprintf(out, "%v %v", i+1, t)
		break
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
