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

	n := readInt()
	cs := make([]string, n)
	ls := make([]int, n)

	for i := range cs {
		cs[i] = readString()
		ls[i] = readInt()
	}

	cnt := 0
	for _, v := range ls {
		cnt += v
		if cnt > 100 {
			fmt.Fprint(out, "Too Long")
			return
		}
	}

	for i, c := range cs {
		for j := 0; j < ls[i]; j++ {
			fmt.Fprint(out, c)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
