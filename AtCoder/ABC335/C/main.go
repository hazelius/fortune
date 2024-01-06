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

	n, q := readInt(), readInt()
	as := make([][]int, n)
	for i := range as {
		as[i] = []int{i + 1, 0}
	}

	head := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			hp := as[head]
			head--
			if head < 0 {
				head = n - 1
			}

			switch readString() {
			case "R":
				as[head] = []int{hp[0] + 1, hp[1]}
			case "L":
				as[head] = []int{hp[0] - 1, hp[1]}
			case "U":
				as[head] = []int{hp[0], hp[1] + 1}
			case "D":
				as[head] = []int{hp[0], hp[1] - 1}
			}
		case 2:
			p := readInt() - 1
			point := as[(p+head)%n]
			fmt.Fprintln(out, point[0], point[1])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
