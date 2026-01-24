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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	q := readInt()
	flg := false
	n := 0
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			n++
		case 2:
			if n > 0 {
				n--
			}
		case 3:
			flg = !flg
		}
		if n >= 3 && flg {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
