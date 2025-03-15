package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	ans := 0
	chars := []rune{'i', 'o'}
	flg := 0
	for _, c := range s {
		if c != chars[flg] {
			ans++
			flg = 1 - flg
		}
		flg = 1 - flg
	}
	if flg > 0 {
		ans++
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
