// https://atcoder.jp/contests/abc243/tasks/abc243_e
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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	abcs := make([][]int, m)
	for i := range abcs {
		abcs[i] = []int{readInt() - 1, readInt() - 1, readInt()}
	}

	ds := make([][]int, n)
	for i := range ds {
		ds[i] = make([]int, n)
	}

	for _, abc := range abcs {
		a, b, c := abc[0], abc[1], abc[2]
		ds[a][b] = c
		ds[b][a] = c
	}

	for i := 0; i < n; i++ {
		for a := 0; a < n; a++ {
			for b := 0; b < n; b++ {
				ai := ds[a][i]
				ib := ds[i][b]
				if ai == 0 || ib == 0 {
					continue
				}
				if ds[a][b] == 0 || ds[a][b] > ai+ib {
					ds[a][b] = ai + ib
				}
			}
		}
	}

	ans := 0
	for _, abc := range abcs {
		a, b, c := abc[0], abc[1], abc[2]
		for i := 0; i < n; i++ {
			ai := ds[a][i]
			ib := ds[i][b]
			if ai > 0 && ib > 0 && c >= ai+ib {
				ans++
				break
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
