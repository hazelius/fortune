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

	n := readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	smap := make(map[string]bool)
	for i, v1 := range ss {
		for j, v2 := range ss {
			if i == j {
				continue
			}
			s := v1 + v2
			smap[s] = true
		}

	}

	fmt.Fprint(out, len(smap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
