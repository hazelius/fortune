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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()

	stat := 0
	cnt := n
	for i, c := range s {
		if c == 'o' {
			stat++
		}
		stat--
		if stat < 0 {
			fmt.Fprintln(out, i+1)
			stat = 0
			cnt--
		}
	}

	for i := 0; i < cnt; i++ {
		fmt.Fprintln(out, n)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
