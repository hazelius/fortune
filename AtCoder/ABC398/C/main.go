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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	m := make(map[int]int)
	for _, a := range as {
		m[a]++
	}

	mx := -1
	for k, v := range m {
		if v > 1 {
			continue
		}
		if k > mx {
			mx = k
		}
	}
	if mx < 0 {
		fmt.Fprint(out, -1)
		return
	}

	for i, v := range as {
		if v == mx {
			fmt.Fprint(out, i+1)
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
