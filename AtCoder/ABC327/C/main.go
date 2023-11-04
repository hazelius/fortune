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

	as := make([][]int, 9)
	for i := range as {
		as[i] = make([]int, 9)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	for _, row := range as {
		rmap := make(map[int]bool)
		for _, v := range row {
			if rmap[v] {
				fmt.Fprint(out, "No")
				return
			}
			rmap[v] = true
		}
	}

	for i := 0; i < 9; i++ {
		cmap := make(map[int]bool)
		for j := 0; j < 9; j++ {
			v := as[j][i]
			if cmap[v] {
				fmt.Fprint(out, "No")
				return
			}
			cmap[v] = true
		}
	}

	qs := [][]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}
	for _, q := range qs {
		smap := make(map[int]bool)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				v := as[i+q[0]][j+q[1]]
				if smap[v] {
					fmt.Fprint(out, "No")
					return
				}
				smap[v] = true
			}
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
