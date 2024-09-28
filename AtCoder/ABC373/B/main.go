package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	s := readString()
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ans := 0
	pre := -1
	for i := range str {
		pos := strings.Index(s, str[i:i+1])
		if pre >= 0 {
			ans += abs(pos - pre)
		}
		pre = pos
	}
	fmt.Fprint(out, ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
