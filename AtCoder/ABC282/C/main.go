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
	sb := []byte(readString())

	flg := true
	for i := 0; i < n; i++ {
		c := sb[i]
		switch c {
		case '"':
			flg = !flg
		case ',':
			if flg {
				sb[i] = byte('.')
			}
		}

	}
	fmt.Fprint(out, string(sb))
}

func main() {
	run(os.Stdin, os.Stdout)
}
