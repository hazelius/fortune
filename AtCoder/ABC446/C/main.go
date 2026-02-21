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
	for cs := 0; cs < t; cs++ {
		n, d := readInt(), readInt()
		as := make([]int, n)
		for i := range as {
			as[i] = readInt()
		}
		bs := make([]int, n)
		for i := range bs {
			bs[i] = readInt()
		}

		total := 0
		used := 0
		didxes := make([]int, n)
		for i, a := range as {
			total += a
			didxes[i] = total

			b := bs[i]
			used += b
			if i-d >= 0 {
				didx := didxes[i-d]
				if didx > used {
					used = didx
				}
			}
		}

		fmt.Fprintln(out, total-used)
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
