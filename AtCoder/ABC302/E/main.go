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

	n, q := readInt(), readInt()

	uvmaps := make(map[int]map[int]bool)

	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			u, v := readInt(), readInt()
			if _, ok := uvmaps[u]; !ok {
				uvmaps[u] = make(map[int]bool)
			}
			uvmaps[u][v] = true

			if _, ok := uvmaps[v]; !ok {
				uvmaps[v] = make(map[int]bool)
			}
			uvmaps[v][u] = true
		case 2:
			v := readInt()
			uvs, ok := uvmaps[v]
			if ok {
				for u := range uvs {
					vm, ok := uvmaps[u]
					if ok {
						delete(vm, v)
						if len(vm) == 0 {
							delete(uvmaps, u)
						}
					}
				}

				delete(uvmaps, v)
			}
		}
		fmt.Fprintln(out, n-len(uvmaps))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
