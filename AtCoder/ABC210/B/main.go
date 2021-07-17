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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	readInt()
	s := readString()
	for i, c := range s {
		if c == '1' {
			if i%2 == 0 {
				return "Takahashi"
			}
			return "Aoki"
		}
	}
	return "Draw"
}

func main() {
	fmt.Println(run(os.Stdin))
}
