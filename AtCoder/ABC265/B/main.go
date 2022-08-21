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

	n, m, t := readInt(), readInt(), readInt()
	as := make([]int, n-1)
	for i := range as {
		as[i] = readInt()
	}
	xymap := make(map[int]int)
	for i := 0; i < m; i++ {
		x, y := readInt(), readInt()
		xymap[x] = y
	}

	for i, a := range as {
		t -= a
		if t <= 0 {
			fmt.Fprint(out, "No")
			return
		}
		y, ok := xymap[i+2]
		if ok {
			t += y
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
