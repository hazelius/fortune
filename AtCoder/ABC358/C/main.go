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

	n, m := readInt(), readInt()
	ss := make([]int, n)
	for i := range ss {
		s := readString()
		for j := range s {
			if s[j] == 'o' {
				ss[i] |= 1 << j
			}
		}
	}

	ans := n
	for i := 0; i < (1 << n); i++ {
		cnt := 0
		v := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				v |= ss[j]
				cnt++
			}
		}
		if v == (1<<m)-1 && ans > cnt {
			ans = cnt
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
