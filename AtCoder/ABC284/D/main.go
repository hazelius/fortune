package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
		n := readInt()
		for i := 2; i*i*i <= n; i++ {
			if n%i > 0 {
				continue
			}
			p := n / i
			if p%i == 0 {
				fmt.Fprintf(out, "%v %v\n", i, p/i)
			} else {
				fmt.Fprintf(out, "%.0f %v\n", math.Sqrt(float64(p)), i)
			}
			break
		}
	}
}
func main() {
	run(os.Stdin, os.Stdout)
}
