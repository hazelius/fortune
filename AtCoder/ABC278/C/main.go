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

type user struct {
	a int
	b int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	_, q := readInt(), readInt()
	nmap := make(map[user]bool)
	for i := 0; i < q; i++ {
		t, a, b := readInt(), readInt(), readInt()
		switch t {
		case 1:
			u := user{a, b}
			nmap[u] = true
		case 2:
			u := user{a, b}
			nmap[u] = false
		case 3:
			if nmap[user{a, b}] && nmap[user{b, a}] {
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
