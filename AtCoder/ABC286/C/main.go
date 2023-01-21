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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a, b := readInt(), readInt(), readInt()
	s := readString()

	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n/2; j++ {
			idx := (n - 1 - j + i) % n
			if s[(j+i)%n] != s[idx] {
				cnt++
			}
		}
		sm := i*a + cnt*b
		if i == 0 || ans > sm {
			ans = sm
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
