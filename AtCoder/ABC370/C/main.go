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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s, t := readString(), readString()
	sbyte, tbyte := []byte(s), []byte(t)
	x := make([]string, len(s))
	cnt := 0
	for i, c := range sbyte {
		if c > tbyte[i] {
			sbyte[i] = tbyte[i]
			x[cnt] = string(sbyte)
			cnt++
		}
	}
	for i := len(sbyte) - 1; i >= 0; i-- {
		if sbyte[i] < tbyte[i] {
			sbyte[i] = tbyte[i]
			x[cnt] = string(sbyte)
			cnt++
		}
	}

	fmt.Fprintln(out, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(out, x[i])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
