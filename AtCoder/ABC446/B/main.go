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

	n, m := readInt(), readInt()
	mmap := make(map[int]bool)
	for i := 1; i <= m; i++ {
		mmap[i] = true
	}

	lxs := make([][]int, n)
	for i := range lxs {
		l := readInt()
		lxs[i] = make([]int, l)
		for j := range lxs[i] {
			lxs[i][j] = readInt()
		}
	}

	for _, lx := range lxs {
		flg := false
		for i := 0; i < len(lx); i++ {
			if mmap[lx[i]] {
				fmt.Fprintln(out, lx[i])
				delete(mmap, lx[i])
				flg = true
				break
			}
		}
		if !flg {
			fmt.Fprintln(out, 0)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
