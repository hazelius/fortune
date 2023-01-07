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
	for i := 0; i < n; i++ {
		n2 := readInt()
		ans := 0
		for j := 0; j < n2; j++ {
			a := readInt()
			if a%2 > 0 {
				ans++
			}
		}

		fmt.Fprintln(out, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
