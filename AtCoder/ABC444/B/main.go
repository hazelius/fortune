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

	n, k := readInt(), readInt()
	ans := 0
	for i := 1; i <= n; i++ {
		tmp := 0
		idx := i
		for idx > 0 {
			tmp += idx % 10
			idx /= 10
		}
		if tmp == k {
			ans++
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
