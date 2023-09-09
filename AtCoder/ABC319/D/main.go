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

	n, m := readInt(), readInt()
	ls := make([]int, n)
	for i := range ls {
		ls[i] = readInt() + 1
	}

	lsum := 0
	for _, v := range ls {
		lsum += v
	}

	ans := sort.Search(lsum, func(i int) bool {
		cnt := 1
		rowLen := 0
		for _, l := range ls {
			if l > i {
				return false
			}
			rowLen += l
			if rowLen > i {
				cnt++
				rowLen = l
			}
		}
		return cnt <= m
	})

	fmt.Fprint(out, ans-1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
