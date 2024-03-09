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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	as := make([]int, 100)
	s := 0
	for i := range as {
		as[i] = readInt()
		if as[i] == 0 {
			s = i
			break
		}
	}
	for i := s; i >= 0; i-- {
		fmt.Fprintln(out, as[i])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
