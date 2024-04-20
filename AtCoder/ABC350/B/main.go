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
	tmap := make(map[int]bool)
	for i := 1; i <= n; i++ {
		tmap[i] = true
	}
	for i := 0; i < q; i++ {
		t := readInt()
		if _, ok := tmap[t]; ok {
			delete(tmap, t)
		} else {
			tmap[t] = true
		}
	}
	fmt.Fprint(out, len(tmap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
