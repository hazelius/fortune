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

	_, q := readInt(), readInt()
	qs := make([][]int, q)
	names := make(map[int]string)
	for i := range qs {
		qs[i] = []int{readInt(), readInt()}
		if qs[i][0] == 2 {
			names[i] = readString()
		}
	}

	idx := 0
	ans := make([]int, 0, q)
	for i := q - 1; i >= 0; i-- {
		p := qs[i][1]
		switch qs[i][0] {
		case 1:
			if p == idx {
				idx = 0
			}
		case 2:
			if p == idx {
				ans = append(ans, i)
			}
		case 3:
			if idx == 0 {
				idx = p
			}
		}
	}

	if len(ans) == 0 {
		fmt.Fprint(out, "")
		return
	}
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(out, names[ans[i]])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
