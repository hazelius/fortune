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
	cs := make([]string, h)
	for i := range cs {
		cs[i] = readString()
	}

	for i := 0; i < w; i++ {
		cnt := 0
		for j := 0; j < h; j++ {
			if cs[j][i] == '#' {
				cnt++
			}
		}
		fmt.Fprint(out, cnt)
		fmt.Fprint(out, " ")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
