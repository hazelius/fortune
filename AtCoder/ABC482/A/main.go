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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	s := readString()
	cnt := 0
	for _, v := range s {
		if v == 'o' {
			cnt++
		}
	}
	if cnt%2 == 1 {
		fmt.Fprint(out, "odd")
	} else {
		fmt.Fprint(out, "even")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
