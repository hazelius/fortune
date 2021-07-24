package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner
var mod int = 1e9 + 7

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()

	chokudai := "chokudai"
	dp := make([]int, len(chokudai))
	for _, c := range s {
		for i, v := range chokudai {
			if v == c {
				if i == 0 {
					dp[0]++
				} else {
					dp[i] = (dp[i] + dp[i-1]) % mod
				}
				break
			}
		}
	}
	return dp[len(chokudai)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
