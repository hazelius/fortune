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

	n := readInt()
	ansmap := make(map[int][]int)
	for i := 0; i < n; i++ {
		k := readInt()
		for j := 0; j < k; j++ {
			a := readInt()
			as, ok := ansmap[a]
			if !ok {
				as = make([]int, 0)
			}
			as = append(as, i+1)
			ansmap[a] = as
		}
	}

	for i := 1; i <= n; i++ {
		as, ok := ansmap[i]
		if !ok {
			fmt.Fprintln(out, 0)
			continue
		}
		fmt.Fprint(out, len((as)))
		for _, a := range as {
			fmt.Fprint(out, " ", a)
		}
		fmt.Fprint(out, "\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
