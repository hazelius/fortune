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
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	mmap := make(map[int]string)
	col := "a"
	clears := make(map[int]bool)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x := readInt()
			if _, ok := clears[x]; ok {
				delete(clears, x)
			} else {
				if _, ok := mmap[x]; !ok {
					mmap[x] = col
				} else {
					clears[x] = true
				}
			}

		case 2:
			col = readString()
			for v := range clears {
				delete(mmap, v)
			}
			clears = make(map[int]bool)
		}
	}

	for i := 1; i <= n; i++ {
		v, ok := mmap[i]
		if ok {
			fmt.Fprint(out, v)
		} else {
			fmt.Fprint(out, col)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
