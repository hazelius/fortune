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

	readInt()
	s := readString()

	abcmap := make(map[rune]bool)
	for i, c := range s {
		abcmap[c] = true
		if len(abcmap) == 3 {
			fmt.Fprint(out, i+1)
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
