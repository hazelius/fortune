package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type point struct {
	x, y int
}

type param struct {
	p   point
	cnt int
}

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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	h, w, k := readInt(), readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}
	hmap := make(map[int]bool)
	wmap := make(map[int]bool)
	for i, s := range ss {
		for j, c := range s {
			if c == '#' {
				hmap[i] = true
				wmap[j] = true
			}
		}
	}

	ans := 0
	queue := make([]param, 0)
	origin := make(map[point]bool)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !hmap[i] && !wmap[j] {
				origin[point{x: i, y: j}] = true
				queue = append(queue, param{p: point{x: i, y: j}, cnt: k})
			}
		}
	}

	dist := []point{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	visited := make(map[point]bool)
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if visited[p.p] {
			continue
		}
		visited[p.p] = true
		ans++

		p.cnt--
		if p.cnt < 0 {
			continue
		}

		for _, d := range dist {
			newP := point{x: p.p.x + d.x, y: p.p.y + d.y}
			if origin[newP] {
				continue
			}
			if newP.x < 0 || newP.x >= h || newP.y < 0 || newP.y >= w {
				continue
			}
			if ss[newP.x][newP.y] == '#' {
				continue
			}
			if !visited[newP] {
				queue = append(queue, param{p: newP, cnt: p.cnt})
			}
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
