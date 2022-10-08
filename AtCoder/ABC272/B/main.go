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
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		k := readInt()
		ks := make([]int, k)
		for j := range ks {
			ks[j] = readInt() - 1
		}
		for _, k2 := range ks {
			for _, k3 := range ks {
				as[k2][k3] = 1
				as[k3][k2] = 1
			}
		}
	}

	for _, a := range as {
		for _, a2 := range a {
			if a2 == 0 {
				fmt.Fprint(out, "No")
				return
			}
		}
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
