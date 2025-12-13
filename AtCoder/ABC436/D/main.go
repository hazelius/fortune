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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	hws := make([]string, h)
	for i := range hws {
		hws[i] = readString()
	}

	warps := make(map[byte][][2]int)
	for i, hw := range hws {
		for j := 0; j < w; j++ {
			if hw[j] == '#' || hw[j] == '.' {
				continue
			}
			warps[hw[j]] = append(warps[hw[j]], [2]int{i, j})
		}
	}

	dist := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := make(map[[2]int]bool)
	visitedWarps := make(map[byte]bool)

	que := [][]int{{0, 0, 0}}
	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		ch, cw, cnt := q[0], q[1], q[2]
		if ch == h-1 && cw == w-1 {
			fmt.Fprint(out, cnt)
			return
		}

		cnt++
		for _, d := range dist {
			nh, nw := ch+d[0], cw+d[1]
			if 0 <= nh && nh < h && 0 <= nw && nw < w {
				if hws[nh][nw] == '#' {
					continue
				}
				if !visited[[2]int{nh, nw}] {
					que = append(que, []int{nh, nw, cnt})
					visited[[2]int{nh, nw}] = true

					if nh == h-1 && nw == w-1 {
						fmt.Fprint(out, cnt)
						return
					}
				}
			}
		}

		c := hws[ch][cw]
		if c == '#' || c == '.' || visitedWarps[c] {
			continue
		}
		visitedWarps[c] = true
		for _, pos := range warps[c] {
			if !visited[pos] {
				que = append(que, []int{pos[0], pos[1], cnt})
				visited[pos] = true
				if pos[0] == h-1 && pos[1] == w-1 {
					fmt.Fprint(out, cnt)
					return
				}
			}
		}

	}
	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
