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

	as := make([]int, 8)
	for i := range as {
		as[i] = readInt()
	}

	for i, a := range as {
		if i > 0 && a < as[i-1] {
			fmt.Fprint(out, "No")
			return
		}
		if !(100 <= a && a <= 675) {
			fmt.Fprint(out, "No")
			return
		}
		if a%25 != 0 {
			fmt.Fprint(out, "No")
			return
		}

	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
