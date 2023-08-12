package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

type question struct {
	t int
	x int
	c string
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	readInt()
	s := []byte(readString())
	q := readInt()

	qs := make([]question, q)
	lstIdx := -1
	for i := range qs {
		qs[i] = question{readInt(), readInt(), readString()}
		if qs[i].t != 1 {
			lstIdx = i
		}
	}

	for i, q := range qs {
		if i == lstIdx {
			tmp := string(s)
			if q.t == 2 {
				tmp = strings.ToLower(tmp)
			} else if q.t == 3 {
				tmp = strings.ToUpper(tmp)
			}
			s = []byte(tmp)
			continue
		}
		if q.t == 1 {
			s[q.x-1] = byte(q.c[0])
		}
	}

	fmt.Fprint(out, string(s))
}

func main() {
	run(os.Stdin, os.Stdout)
}
