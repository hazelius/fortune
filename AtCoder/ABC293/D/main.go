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

	n, m := readInt(), readInt()
	ropes := make([][]int, n)
	for i := range ropes {
		ropes[i] = []int{-1, -1}
	}

	for i := 0; i < m; i++ {
		a, b, c, d := readInt()-1, readString(), readInt()-1, readString()
		idx := 0
		if b == "B" {
			idx = 1
		}
		ropes[a][idx] = c

		idx = 0
		if d == "B" {
			idx = 1
		}
		ropes[c][idx] = a
	}

	circled := 0
	noncircle := 0

	used := make(map[int]bool)
	for i, rope := range ropes {
		if used[i] {
			continue
		}
		used[i] = true

		flgCircle := true
		for dir := 0; dir < 2; dir++ {
			next := rope[dir]
			pre := i
			for {
				if used[next] {
					break
				}
				if next < 0 {
					flgCircle = false
					break
				}
				used[next] = true

				if pre == ropes[next][0] {
					pre = next
					next = ropes[next][1]
				} else {
					pre = next
					next = ropes[next][0]
				}
			}
		}

		if flgCircle {
			circled++
		} else {
			noncircle++
		}
	}

	fmt.Fprintf(out, "%v %v", circled, noncircle)
}

func main() {
	run(os.Stdin, os.Stdout)
}
