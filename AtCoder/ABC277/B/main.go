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

	mapss := make(map[string]int)
	for _, s := range ss {
		_, ok := mapss[s]
		if ok {
			fmt.Fprint(out, "No")
			return
		}
		mapss[s] = 1

		switch s[0] {
		case 'H', 'D', 'C', 'S':
		default:
			fmt.Fprint(out, "No")
			return
		}
		switch s[1] {
		case 'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K':
		default:
			fmt.Fprint(out, "No")
			return
		}
	}

	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
