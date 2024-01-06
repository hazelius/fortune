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
	as := make([][]string, n)
	for i := range as {
		as[i] = make([]string, n)
	}

	as[(n)/2][(n)/2] = "T"
	x, y := 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx := 0
	for i := 1; as[x][y] != "T"; i++ {
		as[x][y] = fmt.Sprint(i)
		dir := dirs[idx]
		newx := x + dir[0]
		newy := y + dir[1]

		if newx < 0 || newy < 0 || newx >= n || newy >= n || as[newx][newy] != "" {
			idx = (idx + 1) % len(dirs)
			dir = dirs[idx]
			newx = x + dir[0]
			newy = y + dir[1]
		}
		x = newx
		y = newy
	}

	for _, a := range as {
		astr := fmt.Sprintf("%v", a)
		fmt.Fprintln(out, astr[1:len(astr)-1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
