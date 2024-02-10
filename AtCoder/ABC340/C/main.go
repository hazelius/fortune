package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

var used map[int]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()

	used = make(map[int]int)
	fmt.Fprint(out, f(n))
}

func f(n int) int {
	if cost, ok := used[n]; ok {
		return cost
	}
	n1 := n / 2
	n2 := n1 + n%2
	ret := n
	if n1 > 1 {
		ret += f(n1)
	}
	if n2 > 1 {
		ret += f(n2)
	}
	used[n] = ret

	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
