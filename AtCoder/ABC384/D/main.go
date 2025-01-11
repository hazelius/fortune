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

	n, s := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	sum := 0
	for _, a := range as {
		sum += a
	}
	s %= sum
	if s == 0 {
		fmt.Fprint(out, "Yes")
		return
	}

	as2 := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		if i != 0 {
			as2[i] = as2[i-1]
		}
		as2[i] += as[i%n]
	}

	mapas := make(map[int]bool)
	mapas[0] = true
	for _, a := range as2 {
		mapas[a] = true
	}

	for _, a := range as2 {
		a -= s
		if a < 0 {
			continue
		}
		if mapas[a] {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
