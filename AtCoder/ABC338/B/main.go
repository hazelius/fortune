package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	smap := make(map[rune]int)
	for _, c := range s {
		smap[c] += 1
	}

	ans := 0
	ansc := 'a'
	for i, v := range smap {
		if v > ans {
			ans = v
			ansc = i
		} else if v == ans {
			if ansc > i {
				ansc = i
			}
		}
	}
	fmt.Fprint(out, string(ansc))
}

func main() {
	run(os.Stdin, os.Stdout)
}
