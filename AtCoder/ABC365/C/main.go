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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sort.Ints(as)

	suma := make([]int, n)
	for i := range suma {
		if i > 0 {
			suma[i] = suma[i-1]
		}
		suma[i] += as[i]
	}

	if suma[n-1] <= m {
		fmt.Fprint(out, "infinite")
		return
	}

	ans := sort.Search(10000000000, func(i int) bool {
		idx := sort.Search(n, func(j int) bool {
			return as[j] > i
		})
		sum := i * (n - idx)
		if idx > 0 {
			sum += suma[idx-1]
		}
		return sum > m
	})

	fmt.Fprint(out, ans-1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
