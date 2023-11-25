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

	n, l, r := readInt(), readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := make([]int, n)
	for idx, a := range as {
		temp := 0
		if a <= l {
			temp = l - a
		} else if r <= a {
			temp = r - a
		}
		x := temp + a
		if x < l || r < x {
			x = -temp + a
		}
		ans[idx] = x
	}

	ansstr := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, ansstr[1:len(ansstr)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
