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

	s, t := readInt(), readInt()
	ans := 0
	for i := 0; i <= s; i++ {
		for j := 0; j <= (s - i); j++ {
			c := i * j
			m := s - i - j
			if c == 0 {
				ans += m + 1
			} else {
				ans += min(t/c, m) + 1
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
