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

	n, m := readInt(), readInt()
	amap := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		amap[a] = append(amap[a], b)
		amap[b] = append(amap[b], a)
	}
	for i := 1; i <= n; i++ {
		ar, ok := amap[i]
		if !ok {
			fmt.Fprintln(out, 0)
			continue
		}
		sort.Ints(ar)
		arstr := fmt.Sprintf("%v", ar)
		fmt.Fprintf(out, "%v %v\n", len(ar), arstr[1:len(arstr)-1])
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
