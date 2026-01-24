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

	n, m := readInt(), readInt()
	abmap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		abmap[a] = append(abmap[a], b)
		abmap[b] = append(abmap[b], a)
	}

	for i := 1; i <= n; i++ {
		cnt, _ := abmap[i]

		ans := n - 1 - len(cnt)
		if ans < 3 {
			fmt.Fprint(out, 0, " ")
		} else {
			ans = ans * (ans - 1) * (ans - 2) / 6
			fmt.Fprint(out, ans, " ")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
