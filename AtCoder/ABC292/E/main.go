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

type vec struct {
	start int
	end   int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	uvmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		uvmap[u] = append(uvmap[u], v)
	}

	henmap := make(map[vec]bool)
	for i := 1; i <= n; i++ {
		vs, ok := uvmap[i]
		if !ok {
			continue
		}

		q := vs
		for len(q) > 0 {
			v := q[0]
			q = q[1:]

			hen := vec{start: i, end: v}
			if i == v || henmap[hen] {
				continue
			}
			henmap[hen] = true

			vs, ok = uvmap[v]
			if !ok {
				continue
			}

			q = append(q, vs...)
		}
	}

	fmt.Fprint(out, len(henmap)-m)
}

func main() {
	run(os.Stdin, os.Stdout)
}
