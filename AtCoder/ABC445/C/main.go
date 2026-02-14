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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt() - 1
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}

	for i := n - 1; i >= 0; i-- {
		if i == as[i] {
			ans[i] = i
		} else {
			ans[i] = ans[as[i]]
		}
	}

	for i := range ans {
		ans[i]++
	}

	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
