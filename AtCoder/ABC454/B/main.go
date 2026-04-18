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

	n, m := readInt(), readInt()
	fmap := make(map[int]int)
	for i := 0; i < n; i++ {
		f := readInt()
		fmap[f]++
	}
	flg := true
	for _, f := range fmap {
		if f > 1 {
			flg = false
			fmt.Fprintln(out, "No")
			break
		}
	}
	if flg {
		fmt.Fprintln(out, "Yes")
	}
	if len(fmap) == m {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
