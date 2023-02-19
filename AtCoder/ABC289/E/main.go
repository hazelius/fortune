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

type point struct {
	u int
	v int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		n, m := readInt(), readInt()
		mc := make([]int, n)
		for i := range mc {
			mc[i] = readInt()
		}
		uvmap := make(map[int][]int)
		for i := 0; i < m; i++ {
			u, v := readInt()-1, readInt()-1
			uvmap[u] = append(uvmap[u], v)
			uvmap[v] = append(uvmap[v], u)
		}

		pmap := make(map[point]int)
		fp := point{0, n - 1}
		pmap[fp] = 0
		q := []point{fp}

		for len(q) > 0 {
			p := q[0]
			q = q[1:]

			for _, u := range uvmap[p.u] {
				for _, v := range uvmap[p.v] {
					if mc[u] == mc[v] {
						continue
					}
					newp := point{u, v}
					_, ok := pmap[newp]
					if ok {
						continue
					}
					pmap[newp] = pmap[p] + 1
					q = append(q, newp)
				}
			}
		}

		ans, ok := pmap[point{n - 1, 0}]
		if ok {
			fmt.Fprintln(out, ans)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
