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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, d := readInt(), readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	used := make(map[point]int)
	que := make([][]int, 0)
	for i, s := range ss {
		for j, c := range s {
			if c == 'H' {
				que = append(que, []int{i, j, d})
			}
		}
	}

	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		p := point{q[0], q[1]}
		f := q[2]
		fu, ok := used[p]
		if ok && fu >= f {
			continue
		}
		used[p] = f

		if f <= 0 {
			continue
		}
		f--

		for _, dir := range dirs {
			x := p.x + dir[0]
			y := p.y + dir[1]
			if x < 0 || y < 0 || x >= h || y >= w {
				continue
			}
			if ss[x][y] == '#' {
				continue
			}
			if ss[x][y] == 'H' {
				continue
			}

			next := point{x, y}
			fu, ok := used[next]
			if ok && fu >= f {
				continue
			}

			que = append(que, []int{x, y, f})
		}
	}

	fmt.Fprint(out, len(used))
}

func main() {
	run(os.Stdin, os.Stdout)
}
