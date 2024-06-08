package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner
var memo = make(map[int][]string)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	if n == 0 {
		fmt.Fprint(out, "#")
		return
	}
	memo = make(map[int][]string)
	ans := grid(n)
	for _, s := range ans {
		fmt.Fprintln(out, s)
	}
}

func grid(n int) []string {
	g, ok := memo[n]
	if ok {
		return g
	}
	if n == 1 {
		return []string{
			"###",
			"#.#",
			"###",
		}
	}
	g = grid(n - 1)
	ret := make([]string, len(g)*3)
	idx := 0
	for i := 0; i < 3; i++ {
		for _, sub := range g {
			if i == 1 {
				ret[idx] = sub + strings.Repeat(".", len(sub)) + sub
			} else {
				ret[idx] = sub + sub + sub
			}
			idx++
		}
	}
	memo[n] = ret
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
