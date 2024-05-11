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

var mod = 100000000

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sort.Ints(as)

	ans := 0
	b := 0
	for i, a := range as {
		ans += a * (n - 1)
		t := mod - a
		ass := as[i+1:]
		p := sort.Search(len(ass), func(j int) bool {
			return t <= ass[j]
		})
		b += len(ass) - p
	}
	ans -= mod * b

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
