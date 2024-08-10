package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	ss := make([]string, n)
	ml := 0
	for i := range ss {
		s := readString()
		if ml < len(s) {
			ml = len(s)
		}
		ss[i] = s
	}

	ans := make([]string, ml)

	for i := 0; i < ml; i++ {
		for _, s := range ss {
			if len(s) > i {
				ans[i] = s[i:i+1] + ans[i]
			} else {
				ans[i] = "*" + ans[i]
			}
		}
	}

	for _, s := range ans {
		fmt.Fprintln(out, strings.TrimRight(s, "*"))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
