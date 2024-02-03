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

	h, w, n := readInt(), readInt(), readInt()
	ss := make([][]byte, h)
	for i := range ss {
		s := ""
		for j := 0; j < w; j++ {
			s += "."
		}
		ss[i] = []byte(s)
	}

	dir := [][]int{{h - 1, 0}, {0, 1}, {1, 0}, {0, w - 1}}
	cur := []int{0, 0, 0}
	for i := 0; i < n; i++ {
		switch ss[cur[0]][cur[1]] {
		case '.':
			ss[cur[0]][cur[1]] = byte('#')
			cur[2] = (cur[2] + 1) % 4
		case '#':
			ss[cur[0]][cur[1]] = byte('.')
			cur[2] = (cur[2] + 3) % 4
		}
		cur[0] = (cur[0] + dir[cur[2]][0]) % h
		cur[1] = (cur[1] + dir[cur[2]][1]) % w
	}

	for _, s := range ss {
		fmt.Fprintln(out, string(s))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
