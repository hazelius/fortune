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
	ms := make(map[int]map[int]bool, n)
	ans := 0
	for i := 0; i < m; i++ {
		u, v := readInt()-1, readInt()-1
		if u == v {
			ans++
			continue
		}
		if u > v {
			u, v = v, u
		}
		if ms[u][v] {
			ans++
			continue
		}
		if _, ok := ms[u]; !ok {
			ms[u] = make(map[int]bool)
		}
		ms[u][v] = true
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
