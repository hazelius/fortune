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

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	dirs := [][]int{
		{1, 1},
		{0, 1},
		{-1, 1},
		{1, 0},
		{-1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	for i, s := range ss {
		for j, c := range s {
			if c != 's' {
				continue
			}

			for _, dir := range dirs {
				x, y := i, j
				ans := [][]int{{x + 1, y + 1}}
				for _, t := range "nuke" {
					x += dir[0]
					y += dir[1]

					if x < 0 || y < 0 || x >= h || y >= w {
						break
					}
					if ss[x][y] != byte(t) {
						break
					}
					ans = append(ans, []int{x + 1, y + 1})
				}

				if len(ans) == 5 {
					for _, v := range ans {
						fmt.Fprintf(out, "%v %v\n", v[0], v[1])
					}
					return
				}
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
