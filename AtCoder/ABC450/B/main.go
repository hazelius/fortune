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
	ass := make([][]int, n-1)
	for i := range ass {
		ass[i] = make([]int, n-1)
		for j := i; j < n-1; j++ {
			ass[i][j] = readInt()
		}
	}

	for s := 0; s < n-2; s++ {
		for e := s + 1; e < n-1; e++ {
			for m := s; m < e; m++ {
				if ass[s][e] > ass[s][m]+ass[m+1][e] {
					fmt.Fprint(out, "Yes")
					return
				}
			}
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
