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

	n, k := readInt(), readInt()
	amap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = true
	}

	for i := 0; i < k; i++ {
		if !amap[i] {
			fmt.Fprint(out, i)
			return
		}
	}

	fmt.Fprint(out, k)
}

func main() {
	run(os.Stdin, os.Stdout)
}
