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
	sc.Split(bufio.ScanWords)

	readInt()
	s := readString()

	flg := false
	for _, c := range s {
		if c == '|' {
			flg = !flg
		} else if c == '*' {
			if flg {
				fmt.Fprint(out, "in")
				return
			}
		}
	}
	fmt.Fprint(out, "out")
}

func main() {
	run(os.Stdin, os.Stdout)
}
