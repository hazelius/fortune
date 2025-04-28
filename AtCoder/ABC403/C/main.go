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

	n, m, q := readInt(), readInt(), readInt()
	nmap := make([]map[int]bool, n)
	for i := range nmap {
		nmap[i] = make(map[int]bool)
	}
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x, y := readInt()-1, readInt()-1
			nmap[x][y] = true
		case 2:
			x := readInt() - 1
			nmap[x][m] = true
		case 3:
			x, y := readInt()-1, readInt()-1
			if nmap[x][m] || nmap[x][y] {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
