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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	hls := make([][]int, n)
	for i := range hls {
		hls[i] = []int{readInt(), readInt()}
	}

	sort.Slice(hls, func(i, j int) bool {
		return hls[i][1] > hls[j][1]
	})
	maxhls := make([][]int, n)
	for i, v := range hls {
		maxhls[i] = v
		if i == 0 {
			continue
		}
		if maxhls[i-1][0] > maxhls[i][0] {
			maxhls[i][0] = maxhls[i-1][0]
		}
	}

	q := readInt()
	for i := 0; i < q; i++ {
		t := readInt()
		hl := sort.Search(n, func(i int) bool {
			return maxhls[i][1] <= t
		})
		fmt.Fprintln(out, maxhls[hl-1][0])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
