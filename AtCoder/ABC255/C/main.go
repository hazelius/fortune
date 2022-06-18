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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x, a, d, n := readInt(), readInt(), readInt(), readInt()
	if d < 0 {
		a = a + d*(n-1)
		d *= -1
	}
	if x <= a {
		return a - x
	}

	x -= a
	maxA := d * (n - 1)
	if x >= maxA {
		return x - maxA
	}
	ans := abs(x % d)
	if ans > abs(d-ans) {
		ans = abs(d - ans)
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
