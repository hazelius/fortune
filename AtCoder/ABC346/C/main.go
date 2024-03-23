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
	mapa := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		if a <= k {
			mapa[a] = true
		}
	}

	ans := k * (k + 1) / 2
	for key := range mapa {
		ans -= key
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
