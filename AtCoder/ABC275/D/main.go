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

var memo map[int]int

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	memo = make(map[int]int)
	ans := f(n)
	fmt.Fprint(out, ans)
}

func f(n int) int {
	if n == 0 {
		return 1
	}
	ret, ok := memo[n]
	if ok {
		return ret
	}
	ret = f(n/2) + f(n/3)
	memo[n] = ret
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
