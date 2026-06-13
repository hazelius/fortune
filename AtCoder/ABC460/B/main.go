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

	t := readInt()
	for i := 0; i < t; i++ {
		x1, y1, r1, x2, y2, r2 := readInt(), readInt(), readInt(), readInt(), readInt(), readInt()
		if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) > (r1+r2)*(r1+r2) {
			fmt.Fprintln(out, "No")
			continue
		}
		if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) < (r1-r2)*(r1-r2) {
			fmt.Fprintln(out, "No")
			continue
		}
		fmt.Fprintln(out, "Yes")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
