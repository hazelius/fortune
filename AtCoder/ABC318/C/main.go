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

	n, d, p := readInt(), readInt(), readInt()
	fs := make([]int, n)
	for i := range fs {
		fs[i] = readInt()
	}

	sum := 0
	for _, v := range fs {
		sum += v
	}

	sort.Sort(sort.Reverse(sort.IntSlice(fs)))

	ans := sum
	for i := 0; i < n+d; i += d {
		sum += p
		for j := i; j < i+d; j++ {
			if j >= len(fs) {
				break
			}
			sum -= fs[j]
		}
		ans = min(ans, sum)
	}

	fmt.Fprint(out, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
