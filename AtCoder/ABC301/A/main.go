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
	t, a := 0, 0
	for _, c := range s {
		if c == 'T' {
			t++
		} else {
			a++
		}
	}
	if a == t {
		if s[len(s)-1:] == "A" {
			fmt.Fprint(out, "T")
		} else {
			fmt.Fprint(out, "A")
		}
	} else if a > t {
		fmt.Fprint(out, "A")
	} else {
		fmt.Fprint(out, "T")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
