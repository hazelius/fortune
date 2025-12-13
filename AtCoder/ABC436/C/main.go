package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type point struct {
	x, y int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	_, m := readInt(), readInt()
	rcs := make([][]int, m)
	for i := range rcs {
		rcs[i] = []int{readInt(), readInt()}
	}
	nn := make(map[point]bool)

	ans := 0
	for _, rc := range rcs {
		r, c := rc[0], rc[1]
		if !nn[point{r - 1, c - 1}] && !nn[point{r, c}] && !nn[point{r - 1, c}] && !nn[point{r, c - 1}] {
			ans++
			nn[point{r - 1, c - 1}] = true
			nn[point{r, c}] = true
			nn[point{r - 1, c}] = true
			nn[point{r, c - 1}] = true
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
