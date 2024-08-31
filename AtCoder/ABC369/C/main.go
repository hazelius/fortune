package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	ans := 0
	l := 0
	for r := 0; r < n; r++ {
		if l == r {
			ans++
			continue
		}
		if l == r-1 {
			ans += 2
			continue
		}
		if as[r]-as[r-1] != as[r-1]-as[r-2] {
			ans += 2
			l = r - 1
			continue
		}
		ans += r - l + 1
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
