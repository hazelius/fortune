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
	aaa := make([][][]int, n)
	for i := range aaa {
		aaa[i] = make([][]int, n)
		for j := range aaa[i] {
			aaa[i][j] = make([]int, n)
			for k := range aaa[i][j] {
				a := readInt()
				if i > 0 {
					a += aaa[i-1][j][k]
				}
				if j > 0 {
					a += aaa[i][j-1][k]
				}
				if k > 0 {
					a += aaa[i][j][k-1]
				}
				if i > 0 && j > 0 {
					a -= aaa[i-1][j-1][k]
				}
				if i > 0 && k > 0 {
					a -= aaa[i-1][j][k-1]
				}
				if j > 0 && k > 0 {
					a -= aaa[i][j-1][k-1]
				}
				if i > 0 && j > 0 && k > 0 {
					a += aaa[i-1][j-1][k-1]
				}
				aaa[i][j][k] = a
			}
		}
	}

	q := readInt()
	for i := 0; i < q; i++ {
		lx, rx, ly, ry, lz, rz := readInt()-2, readInt()-1, readInt()-2, readInt()-1, readInt()-2, readInt()-1
		ans := aaa[rx][ry][rz]
		if lz >= 0 {
			ans -= aaa[rx][ry][lz]
		}
		if ly >= 0 {
			ans -= aaa[rx][ly][rz]
		}
		if lx >= 0 {
			ans -= aaa[lx][ry][rz]
		}
		if lz >= 0 && ly >= 0 {
			ans += aaa[rx][ly][lz]
		}
		if lz >= 0 && lx >= 0 {
			ans += aaa[lx][ry][lz]
		}
		if ly >= 0 && lx >= 0 {
			ans += aaa[lx][ly][rz]
		}
		if lz >= 0 && ly >= 0 && lx >= 0 {
			ans -= aaa[lx][ly][lz]
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
