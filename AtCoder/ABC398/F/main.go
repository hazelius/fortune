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
	ln := len(s)
	for i := 0; i < ln; i++ {
		flg := true
		for j := i; j < ln; j++ {
			if s[j] != s[ln-1-j+i] {
				flg = false
				break
			}
		}
		if !flg {
			continue
		}

		fmt.Fprint(out, s)
		for j := i - 1; j >= 0; j-- {
			fmt.Fprint(out, string(s[j]))
		}
		return
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
