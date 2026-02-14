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
	ss := make([]string, n)
	lst := 0
	for i := range ss {
		ss[i] = readString()
		if len(ss[i]) > lst {
			lst = len(ss[i])
		}
	}

	for _, s := range ss {
		cnt := (lst - len(s)) / 2
		for i := 0; i < cnt; i++ {
			fmt.Fprint(out, ".")
		}
		fmt.Fprint(out, s)
		for i := 0; i < cnt; i++ {
			fmt.Fprint(out, ".")
		}
		fmt.Fprintln(out, "")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
