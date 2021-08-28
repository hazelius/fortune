// B - Union Find
// https://atcoder.jp/contests/atc001/tasks/unionfind_a
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	n, q := readInt(), readInt()
	ts := make([]int, n+1)
	for i := range ts {
		ts[i] = i
	}

	ans := make([]string, 0)
	for i := 0; i < q; i++ {
		p, a, b := readInt(), readInt(), readInt()
		if p == 0 {
			unite(a, b, ts)
		} else {
			if same(a, b, ts) {
				ans = append(ans, "Yes")
			} else {
				ans = append(ans, "No")
			}
		}
	}

	return strings.Join(ans, "\n")
}

func root(i int, ts []int) int {
	if ts[i] == i {
		return i
	}
	ts[i] = root(ts[i], ts)
	return ts[i]
}

func same(x, y int, ts []int) bool {
	return (root(x, ts) == root(y, ts))
}

func unite(x, y int, ts []int) {
	x = root(x, ts)
	y = root(y, ts)
	if x == y {
		return
	}
	ts[x] = y
}

func main() {
	fmt.Println(run(os.Stdin))
}
