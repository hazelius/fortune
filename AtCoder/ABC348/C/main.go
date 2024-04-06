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
	acmap := make(map[int]int)
	for i := 0; i < n; i++ {
		a, c := readInt(), readInt()
		v, ok := acmap[c]
		if !ok {
			acmap[c] = a
		} else if a < v {
			acmap[c] = a
		}
	}
	ans := 0
	for _, v := range acmap {
		if v > ans {
			ans = v
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
