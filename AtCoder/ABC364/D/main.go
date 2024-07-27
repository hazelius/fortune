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

	n, q := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sort.Ints(as)

	for i := 0; i < q; i++ {
		b, k := readInt(), readInt()

		ans := sort.Search(1000000000, func(v int) bool {
			l := sort.Search(n, func(w int) bool {
				return as[w] >= b-v
			})
			r := sort.Search(n, func(w int) bool {
				return as[w] > b+v
			})
			r--
			return r-l+1 >= k
		})

		fmt.Fprintln(out, ans)
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
