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

	n := readInt()
	vmap := make(map[int]bool)
	abmap := make(map[int][]int)
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		if a == 0 && b == 0 {
			vmap[i+1] = true
			continue
		}
		abmap[a] = append(abmap[a], i+1)
		abmap[b] = append(abmap[b], i+1)
	}

	que := make([]int, 0)
	for k := range vmap {
		que = append(que, k)
	}

	for len(que) > 0 {
		q := que[0]
		que = que[1:]

		abs, ok := abmap[q]
		if !ok {
			continue
		}

		for _, v := range abs {
			if !vmap[v] {
				vmap[v] = true
				que = append(que, v)
			}
		}
	}

	fmt.Fprint(out, len(vmap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
