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
	sc.Split(bufio.ScanWords)

	n := readInt()
	wxs := make([][]int, n)
	for i := range wxs {
		wxs[i] = []int{readInt(), readInt()}
	}

	ans := 0
	for i := 0; i < 24; i++ {
		cnt := 0
		for _, wx := range wxs {
			w, x := wx[0], wx[1]
			t := (x + i) % 24
			if t >= 9 && t < 18 {
				cnt += w
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
