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
	sc.Split(bufio.ScanWords)

	h, w := readInt(), readInt()
	as := make([]string, h)
	for i := range as {
		as[i] = readString()
	}
	bs := make([]string, h)
	for i := range bs {
		bs[i] = readString()
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			flg := true
			for bi, b := range bs {
				idx := (bi + j) % h
				a := as[idx][i:] + as[idx][:i]
				if b != a {
					flg = false
					break
				}
			}
			if flg {
				fmt.Fprint(out, "Yes")
				return
			}
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
