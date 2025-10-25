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
	h int
	w int
}

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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	pmap := make(map[point]bool)
	for i, s := range ss {
		for j, c := range s {
			if c == '#' {
				pmap[point{i, j}] = true
			}
		}
	}

	que := pmap
	for len(que) > 0 {
		que = f(h, w, que, pmap)
		for tar := range que {
			pmap[tar] = true
		}
	}

	fmt.Fprint(out, len(pmap))
}

func f(h, w int, que map[point]bool, pmap map[point]bool) map[point]bool {
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	ret := make(map[point]bool)
	for q := range que {
		for _, d1 := range dir {
			tar := point{q.h + d1[0], q.w + d1[1]}
			if tar.h < 0 || tar.w < 0 || tar.h >= h || tar.w >= w {
				continue
			}
			if pmap[tar] || ret[tar] {
				continue
			}

			cnt := 0
			for _, d2 := range dir {
				d := point{tar.h + d2[0], tar.w + d2[1]}
				if d.h < 0 || d.w < 0 || d.h >= h || d.w >= w {
					continue
				}

				if pmap[d] {
					cnt++
				}
			}
			if cnt != 1 {
				continue
			}
			ret[tar] = true
		}
	}

	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
