package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type point struct {
	x int
	y int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	_, m := readInt(), readInt()
	xmap := make(map[int][]point)
	ymap := make(map[int][]point)
	for i := 0; i < m; i++ {
		rc := point{readInt(), readInt()}
		delete(xmap, rc.x)
		delete(ymap, rc.y)
		xmap[rc.x] = append(xmap[rc.x], rc)
		ymap[rc.y] = append(ymap[rc.y], rc)
	}
	ysmap := make(map[point]bool)
	for _, rcs := range ymap {
		for _, rc := range rcs {
			ysmap[rc] = true
		}
	}

	ans := 0
	for _, rcs := range xmap {
		for _, rc := range rcs {
			if ysmap[rc] {
				ans++
			}
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
