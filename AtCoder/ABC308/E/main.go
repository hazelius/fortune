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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	s := readString()

	mix := make([][]int, n)
	for i := 0; i < n; i++ {
		mix[i] = make([]int, 3)
		if i > 0 {
			copy(mix[i], mix[i-1])
		}
		if s[i] == 'M' {
			mix[i][as[i]]++
		}
	}

	xix := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		xix[i] = make([]int, 3)
		if i < n-1 {
			copy(xix[i], xix[i+1])
		}
		if s[i] == 'X' {
			xix[i][as[i]]++
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		if s[i] != 'E' {
			continue
		}

		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ans += mix[i-1][j] * xix[i+1][k] * mex(j, k, as[i])
			}
		}
	}

	fmt.Fprint(out, ans)
}

func mex(a, b, c int) int {
	for i := 0; i < 3; i++ {
		if a != i && b != i && c != i {
			return i
		}
	}
	return 3
}

func main() {
	run(os.Stdin, os.Stdout)
}
