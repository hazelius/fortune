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
	ps := make([]int, n)
	for i := range ps {
		ps[i] = readInt()
	}

	price := sort.Search(m, func(i int) bool {
		sum := 0
		for _, p := range ps {
			k := ((i-1)/p + 1) / 2
			if k <= 0 {
				continue
			}
			if m/k/k/p == 0 {
				return true
			}
			sum += k * k * p
			if sum > m {
				return true
			}
		}
		return false
	})

	price--
	sum := 0
	ans := 0
	for _, p := range ps {
		k := ((price-1)/p + 1) / 2
		if k <= 0 {
			continue
		}
		ans += k
		sum += k * k * p
	}

	ans += (m - sum) / (price + 1)
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
