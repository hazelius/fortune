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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := make([]int, n)
	minus := make([]int, n)
	give := n
	take := -1
	for i, a := range as {
		give--
		take = take + 1 - minus[i]
		a = a - give + take
		if a < 0 {
			minus[n+a]++
			a = 0
		}
		ans[i] = a
	}

	str := fmt.Sprintf("%d", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
