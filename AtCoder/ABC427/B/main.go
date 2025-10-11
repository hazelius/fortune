package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var memo map[int]int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	memo = make(map[int]int)
	memo[0] = 1
	ans := 0
	for i := 0; i < n; i++ {
		m := memo[i]
		ret := f(m)
		ans += ret
		memo[i+1] = ans
	}
	fmt.Fprint(out, ans)
}

func f(a int) int {
	ret := 0
	for {
		ret += a % 10
		if a < 10 {
			break
		}
		a /= 10
	}
	return ret
}

func main() {
	run(os.Stdin, os.Stdout)
}
