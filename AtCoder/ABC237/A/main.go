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

	n := readString()
	if len(n) > 15 {
		return "No"
	}
	nint, _ := strconv.Atoi(n)
	if nint < 0 {
		nint *= -1
		nint--
	}

	i := 0
	for s := 1; s <= nint; s *= 2 {
		i++
		if s >= nint {
			return "No"
		}
	}
	if i > 31 {
		return "No"
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
