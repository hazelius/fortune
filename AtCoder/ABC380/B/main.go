package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	ss := strings.Split(s, "|")
	ss = ss[1 : len(ss)-1]

	ans := make([]int, len(ss))
	for i, v := range ss {
		ans[i] = len(v)
	}
	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
