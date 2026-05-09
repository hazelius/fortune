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

	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := sort.Search(as[0]+k+1, func(x int) bool {
		cnt := 0
		for i, a := range as {
			if a < x {
				cnt += (x - a + i) / (i + 1)
			}
			if cnt > k {
				return true
			}
		}
		return cnt > k
	})

	fmt.Fprint(out, ans-1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
