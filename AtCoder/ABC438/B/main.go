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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	s := readString()
	t := readString()
	si := make([]int, n)
	ti := make([]int, m)
	for i := range si {
		si[i], _ = strconv.Atoi(string(s[i]))
	}
	for i := range ti {
		ti[i], _ = strconv.Atoi(string(t[i]))
	}

	ans := -1
	for i := 0; i <= n-m; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			cnt += (10 + si[i+j] - ti[j]) % 10
		}
		if ans < 0 || cnt < ans {
			ans = cnt
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
