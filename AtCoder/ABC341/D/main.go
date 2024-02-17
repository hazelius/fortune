package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	n, m, k := readInt(), readInt(), readInt()

	n2, m2 := n, m
	kb := 0
	for {
		if n2 < m2 {
			n2 += n
		} else if n2 == m2 {
			kb = n2
			break
		} else {
			m2 += m
		}
	}

	ans := sort.Search(math.MaxInt64, func(j int) bool {
		cnt := (j / n) + (j / m) - (j/kb)*2
		return cnt >= k
	})
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
