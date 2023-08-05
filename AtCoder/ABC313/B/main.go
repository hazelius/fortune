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

	n, m := readInt(), readInt()
	abs := make([]int, n)
	for i := range abs {
		abs[i] = -1
	}
	for i := 0; i < m; i++ {
		a, b := readInt()-1, readInt()-1
		abs[b] = a
	}

	for flg := true; flg; {
		flg = false
		for b, a := range abs {
			if a != -1 && a != abs[a] && abs[a] != -1 {
				abs[b] = abs[a]
				flg = true
			}
		}
	}

	ans := 0
	flg := false
	for i, v := range abs {
		if v != -1 {
			continue
		}
		if flg {
			fmt.Fprint(out, -1)
			return
		}
		flg = true
		ans = i
	}

	fmt.Fprint(out, ans+1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
