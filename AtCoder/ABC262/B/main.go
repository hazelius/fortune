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

	n, m := readInt(), readInt()
	uvmap := make(map[int]map[int]bool)
	for i := 0; i < m; i++ {
		u, v := readInt(), readInt()
		if _, ok := uvmap[u]; !ok {
			uvmap[u] = make(map[int]bool)
		}
		if _, ok := uvmap[v]; !ok {
			uvmap[v] = make(map[int]bool)
		}
		uvmap[u][v] = true
		uvmap[v][u] = true
	}

	ans := 0
	for i := 1; i <= n; i++ {
		vs, ok := uvmap[i]
		if !ok {
			continue
		}
		for k := range vs {
			if k < i {
				continue
			}
			ks, ok := uvmap[k]
			if !ok {
				continue
			}
			for k2 := range ks {
				if k2 < k {
					continue
				}
				if uvmap[k2][i] {
					ans++
				}
			}
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
