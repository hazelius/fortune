package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	s := k / n
	a := k % n

	asort := make([]int, n)
	copy(asort, as)
	sort.Ints(asort)
	last := asort[a]

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		v := s

		if as[i] < last {
			v++
		}
		ans[i] = strconv.Itoa(v)
	}
	return strings.Join(ans, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
