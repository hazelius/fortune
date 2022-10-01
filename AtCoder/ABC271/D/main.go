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

	n, s := readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}

	dp := make([]map[int]string, n)
	for i, ab := range abs {
		dp[i] = make(map[int]string)
		h := ab[0]
		t := ab[1]
		if i > 0 {
			for k, v := range dp[i-1] {
				dp[i][k+h] = v + "H"
				dp[i][k+t] = v + "T"
			}
		} else {
			dp[i][h] = "H"
			dp[i][t] = "T"
		}
	}

	ans, ok := dp[n-1][s]
	if !ok {
		fmt.Fprint(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
