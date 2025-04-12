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
	ans := 0
	an := 2
	for a := 1; an <= n; a++ {
		an *= 2
		x := n
		for i := 0; i < a; i++ {
			x /= 2
		}
		m := sort.Search(x, func(i int) bool {
			if i == 0 {
				return false
			}
			return i > x/i
		})
		if m == 1 {
			m++
		}

		ans += m / 2
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
