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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	m := readInt()
	bmap := make(map[int]bool)
	for i := 0; i < m; i++ {
		b := readInt()
		bmap[b] = true
	}
	x := readInt()

	que := []int{0}
	used := make(map[int]bool)
	for len(que) > 0 {
		q := que[0]
		que = que[1:]

		for _, a := range as {
			t := q + a
			if used[t] {
				continue
			}
			if bmap[t] {
				continue
			}
			if t == x {
				fmt.Fprint(out, "Yes")
				return
			}
			if t > x {
				continue
			}
			que = append(que, t)
			used[t] = true
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
