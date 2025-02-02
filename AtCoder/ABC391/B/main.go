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

	n, m := readInt(), readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}
	ts := make([]string, m)
	for i := range ts {
		ts[i] = readString()
	}

	for i := 0; i <= n-m; i++ {
		for j := 0; j <= n-m; j++ {
			if isMatch(i, j, ss, ts) {
				fmt.Fprintf(out, "%v %v", i+1, j+1)
				return
			}
		}
	}

	fmt.Fprint(out, "No")
}

func isMatch(a, b int, ss, ts []string) bool {
	for i, t := range ts {
		if t != ss[a+i][b:b+len(t)] {
			return false
		}
	}
	return true
}

func main() {
	run(os.Stdin, os.Stdout)
}
