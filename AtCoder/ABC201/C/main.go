package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s string) int {
	ans := 0
	for i := 0; i < 10000; i++ {
		if isOK(i, s) {
			ans++
		}
	}
	return ans
}

func isOK(n int, s string) bool {
	for i, c := range s {
		if !sub(i, n, c) {
			return false
		}
	}
	return true
}

func sub(i, n int, c rune) bool {
	for j := 1; j <= 4; j++ {
		v := n % 10
		n /= 10
		if i == v {
			if c == 'x' {
				return false
			} else if c == 'o' {
				return true
			}
		}
	}
	if c == 'o' {
		return false
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
