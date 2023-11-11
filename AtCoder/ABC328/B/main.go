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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ds := make([]int, n)
	for i := range ds {
		ds[i] = readInt()
	}

	ans := 0
	for i, d := range ds {
		m := i + 1
		mc := m % 10
		flg := true
		for m > 9 {
			m /= 10
			tc := m % 10
			if tc != mc {
				flg = false
				break
			}
		}
		if !flg {
			continue
		}

		for d >= mc {
			ans++
			mc = 10*mc + mc
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
