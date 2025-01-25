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

	cnts := make([]int, 0)
	for i := 1; i <= 5; i++ {
		if readInt() != i {
			cnts = append(cnts, i)
		}
	}

	if len(cnts) != 2 || cnts[1]-cnts[0] != 1 {
		fmt.Fprint(out, "No")
	} else {
		fmt.Fprint(out, "Yes")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
