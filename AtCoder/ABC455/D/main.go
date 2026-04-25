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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	cpTo := make([]int, n)
	for i := range cpTo {
		cpTo[i] = -1
	}
	cpFrom := make([]int, n)
	for i := range cpFrom {
		cpFrom[i] = -1
	}

	for i := 0; i < q; i++ {
		c, p := readInt()-1, readInt()-1
		cpTo[p] = c
		if cpFrom[c] != -1 {
			cpTo[cpFrom[c]] = -1
		}
		cpFrom[c] = p
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if cpFrom[i] >= 0 {
			continue
		}
		cnt := 1
		c := i
		for cpTo[c] != -1 {
			cnt++
			c = cpTo[c]
		}
		ans[i] = cnt
	}
	str := fmt.Sprintf("%d", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
