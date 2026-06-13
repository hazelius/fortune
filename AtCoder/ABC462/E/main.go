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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		a, b, x, y := readInt(), readInt(), abs(readInt()), abs(readInt())
		sa := abs(x - y)
		ans := min(x, y) * min(a, b) * 2

		if a*3 < b {
			b = a * 3
		} else if a > b*3 {
			a = b * 3
		}

		ans += (sa/2)*a + (sa/2)*b
		if sa%2 > 0 {
			if x > y {
				ans += a
			} else {
				ans += b
			}
		}

		fmt.Fprintln(out, ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
