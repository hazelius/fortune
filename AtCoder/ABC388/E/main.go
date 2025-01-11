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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := sort.Search((n/2)+1, func(cnt int) bool {
		return !f(cnt, as)
	})
	fmt.Fprint(out, ans-1)

}

func f(cnt int, as []int) bool {
	for i := 0; i < cnt; i++ {
		if as[i]*2 > as[len(as)-cnt+i] {
			return false
		}
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
