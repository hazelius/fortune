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
	cs := make([]string, h)
	for i := range cs {
		cs[i] = readString()
	}
	used := make([][]bool, h)
	for i := range used {
		used[i] = make([]bool, w)
	}

	ans := make([]int, min(h, w))
	for i, c := range cs {
		for j, cc := range c {
			if cc == '#' && !used[i][j] {
				a := f(i, j, cs, used)
				ans[a-1]++
			}
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func f(i, j int, cs []string, used [][]bool) int {
	cnt := 0
	x, y := i, j
	for cs[x][y] == '#' {
		used[x][y] = true
		cnt++
		x++
		y++
		if x >= len(cs) || y >= len(cs[0]) {
			break
		}
	}

	x = i + cnt - 1
	y = j
	for cs[x][y] == '#' {
		used[x][y] = true
		x--
		y++
		if x < 0 || y >= len(cs[0]) {
			break
		}
	}
	return cnt / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
