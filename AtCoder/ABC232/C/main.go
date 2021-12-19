// https://atcoder.jp/contests/abc232/tasks/abc232_c
// C - Graph Isomorphism
// 順列の全列挙

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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	ab := make([][]bool, n)
	for i := range ab {
		ab[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		ab[a][b] = true
		ab[b][a] = true
	}

	cd := make([][]bool, n)
	for i := range cd {
		cd[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		c, d := readInt()-1, readInt()-1
		cd[c][d] = true
		cd[d][c] = true
	}

	p := make([]int, n)
	for i := range p {
		p[i] = i
	}

	for {
		if f(ab, cd, p) {
			return "Yes"
		}

		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}

	return "No"
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func f(ab, cd [][]bool, p []int) bool {
	for i := 0; i < len(ab); i++ {
		for j := 0; j < len(ab[i]); j++ {
			if ab[i][j] != cd[p[i]][p[j]] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(run(os.Stdin))
}
