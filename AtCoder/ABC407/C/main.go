package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	pre := 0
	ans := 0
	n := 0
	for i, c := range s {
		ans++
		n, _ = strconv.Atoi(string(c))
		if i == 0 {
			pre = n
			continue
		}
		if n > pre {
			ans += pre + 10 - n
		} else {
			ans += pre - n
		}
		pre = n
	}
	ans += n
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
