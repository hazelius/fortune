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
	ls := make([]int, n)
	for i := range ls {
		ls[i] = readInt()
	}
	ans := 1
	for _, l := range ls {
		if l == 1 {
			break
		}
		ans++
	}

	ans++
	for i := n - 1; i >= 0; i-- {
		if ls[i] == 1 {
			break
		}
		ans++
	}

	ans = n + 1 - ans
	if ans < 0 {
		ans = 0
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
