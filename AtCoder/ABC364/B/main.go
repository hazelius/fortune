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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	cur := []int{readInt() - 1, readInt() - 1}

	cs := make([]string, h)
	for i := range cs {
		cs[i] = readString()
	}
	x := readString()

	dir := map[rune][]int{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}}
	for _, c := range x {
		d := dir[c]
		next := []int{cur[0] + d[0], cur[1] + d[1]}
		if next[0] < 0 || next[0] >= h || next[1] < 0 || next[1] >= w {
			continue
		}
		if cs[next[0]][next[1]] == '#' {
			continue
		}
		cur = next
	}
	fmt.Fprint(out, cur[0]+1, cur[1]+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
