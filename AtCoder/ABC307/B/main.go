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
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for _, v := range []string{ss[i] + ss[j], ss[j] + ss[i]} {
				if v == Reverse(v) {
					fmt.Fprint(out, "Yes")
					return
				}
			}
		}
	}

	fmt.Fprint(out, "No")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	run(os.Stdin, os.Stdout)
}
