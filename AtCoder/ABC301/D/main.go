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

	s := readString()
	n := readInt()

	ans := 0
	for i, c := range s {
		idx := len(s) - i - 1
		if c == '1' {
			ans += 1 << idx
		}
	}
	if ans > n {
		fmt.Fprint(out, "-1")
		return
	}

	for i, c := range s {
		if c != '?' {
			continue
		}
		idx := len(s) - i - 1
		add := 1 << idx
		if ans+add <= n {
			ans += add
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
