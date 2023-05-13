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

	tmp := as[0]
	idx := 1
	fmt.Fprint(out, tmp)
	for {
		a := as[idx]
		if tmp < a {
			tmp++
			fmt.Fprintf(out, " %v", tmp)
		} else if tmp > a {
			tmp--
			fmt.Fprintf(out, " %v", tmp)
		} else {
			idx++
			if idx >= n {
				break
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
