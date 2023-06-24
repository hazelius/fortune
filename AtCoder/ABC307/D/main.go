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
	s := readString()

	idxs := make([]int, 0, n)
	ans := make([]rune, 0, n)

	for _, c := range s {
		if c == ')' {
			if len(idxs) > 0 {
				idx := idxs[len(idxs)-1]
				ans = ans[:idx]
				idxs = idxs[:len(idxs)-1]
				continue
			}
		} else if c == '(' {
			idxs = append(idxs, len(ans))
		}
		ans = append(ans, c)
	}

	fmt.Fprint(out, string(ans))
}

func main() {
	run(os.Stdin, os.Stdout)
}
