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
	s := readString()

	dp := make([]map[rune]int, n+1)
	dp[0] = make(map[rune]int)
	for i, c := range s {
		dp[i+1] = make(map[rune]int)
		v, e := 'R', c
		switch c {
		case 'R':
			v = 'P'
		case 'P':
			v = 'S'
		}
		f(i, 1, v, dp)
		f(i, 0, e, dp)
	}

	ans := 0
	for _, v := range dp[n] {
		ans = max(ans, v)
	}
	fmt.Fprint(out, ans)
}

func f(i, cnt int, c rune, dp []map[rune]int) {
	v := 0
	for pchar, pcnt := range dp[i] {
		if pchar == c {
			continue
		}
		v = max(v, pcnt)
	}
	v += cnt
	dp[i+1][c] = v
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
