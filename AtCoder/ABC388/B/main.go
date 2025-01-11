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

	n, d := readInt(), readInt()
	tls := make([][]int, n)
	for i := range tls {
		tls[i] = []int{readInt(), readInt()}
	}

	for k := 1; k <= d; k++ {
		sort.Slice(tls, func(i, j int) bool {
			return tls[i][0]*(tls[i][1]+k) > tls[j][0]*(tls[j][1]+k)
		})
		fmt.Fprintln(out, tls[0][0]*(tls[0][1]+k))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
