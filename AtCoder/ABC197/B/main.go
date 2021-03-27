package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(h, w, x, y int, ss []string) int {
	x--
	y--

	if ss[x][y:y+1] == "#" {
		return 0
	}
	r := 1

	for i := x + 1; i < h; i++ {
		if ss[i][y:y+1] == "#" {
			break
		}
		r++
	}

	for i := x - 1; i >= 0; i-- {
		if ss[i][y:y+1] == "#" {
			break
		}
		r++
	}

	for i := y + 1; i < w; i++ {
		if ss[x][i:i+1] == "#" {
			break
		}
		r++
	}

	for i := y - 1; i >= 0; i-- {
		if ss[x][i:i+1] == "#" {
			break
		}
		r++
	}

	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w, x, y := readInt(), readInt(), readInt(), readInt()
	sarr := make([]string, h)
	for i := 0; i < h; i++ {
		sarr[i] = readString()
	}
	fmt.Println(run(h, w, x, y, sarr))
}
