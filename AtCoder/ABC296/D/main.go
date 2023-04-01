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

	n, m := readInt(), readInt()
	ans := -1
	for a := 1; a <= n; a++ {
		b := m / a
		if m%a > 0 {
			b++
		}
		if b <= n {
			if ans == -1 || ans > a*b {
				ans = a * b
			}
		}
		if a > b {
			break
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
