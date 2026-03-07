package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	as := make([][]int, n)
	for i := range as {
		as[i] = []int{i, readInt()}
	}
	sort.Slice(as, func(i, j int) bool {
		return as[i][1] < as[j][1]
	})

	for i := 0; i < q; i++ {
		k := readInt()
		bmap := make(map[int]bool)
		for j := 0; j < k; j++ {
			b := readInt() - 1
			bmap[b] = true
		}
		for _, v := range as {
			if !bmap[v[0]] {
				fmt.Fprintln(out, v[1])
				break
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
