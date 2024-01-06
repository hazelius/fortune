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
	for x := 0; x <= n; x++ {
		for y := 0; y <= n; y++ {
			for z := 0; z <= n; z++ {
				if x+y+z <= n {
					fmt.Fprintln(out, x, y, z)
				}
			}
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
