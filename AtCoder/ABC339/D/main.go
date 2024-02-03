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

type position struct {
	p1x int
	p1y int
	p2x int
	p2y int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	cur := position{}
	flg := false
	for x, s := range ss {
		for y, c := range s {
			if c == 'P' {
				if flg {
					cur.p2x = x
					cur.p2y = y
					break
				}

				cur.p1x = x
				cur.p1y = y
				flg = true
			}
		}
	}

	used := make(map[position]int)
	used[cur] = 0
	dirs := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	que := []position{cur}
	for i := 0; len(que) > 0; i++ {
		pos := que[0]
		que = que[1:]
		cnt := used[pos]
		cnt++

		for _, dir := range dirs {
			new1 := move(pos.p1x, pos.p1y, dir, ss)
			new2 := move(pos.p2x, pos.p2y, dir, ss)
			if new1[0] == new2[0] && new1[1] == new2[1] {
				fmt.Fprint(out, cnt)
				return
			}
			newpos := position{new1[0], new1[1], new2[0], new2[1]}
			if _, ok := used[newpos]; !ok {
				used[newpos] = cnt
				que = append(que, newpos)
			}
		}
	}

	fmt.Fprint(out, -1)
}

func move(x, y int, dir []int, ss []string) []int {
	newx, newy := x+dir[0], y+dir[1]
	if newx < 0 || newy < 0 || newx >= len(ss) || newy >= len(ss[0]) {
		return []int{x, y}
	}
	if ss[newx][newy] == '#' {
		return []int{x, y}
	}

	return []int{newx, newy}
}

func main() {
	run(os.Stdin, os.Stdout)
}
