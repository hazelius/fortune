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

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	for i, s := range ss {
		for j, c := range s {
			if c != '.' {
				continue
			}

			cnt := 0
			if i > 0 && ss[i-1][j] == '#' {
				cnt++
			}
			if i < h-1 && ss[i+1][j] == '#' {
				cnt++
			}
			if j > 0 && ss[i][j-1] == '#' {
				cnt++
			}
			if j < w-1 && ss[i][j+1] == '#' {
				cnt++
			}

			if cnt > 1 {
				fmt.Fprintf(out, "%v %v", i+1, j+1)
				return
			}
		}
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
