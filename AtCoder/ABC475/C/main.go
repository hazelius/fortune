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

	n, s, l := readInt(), readInt()-1, readInt()
	as := make([]int, n-1)
	for i := range as {
		as[i] = readInt()
	}
	sa := make([]int, n)
	for i := s + 1; i < n; i++ {
		sa[i] = sa[i-1] + as[i-1]
	}
	for i := s - 1; i >= 0; i-- {
		sa[i] = sa[i+1] + as[i]
	}

	ans := 0
	for i, v := range sa {
		if i > s {
			break
		}
		if v > l {
			continue
		}
		tmp := s - i + 1
		for j := s + 1; j < n; j++ {
			if v+sa[j]*2 <= l || v*2+sa[j] <= l {
				tmp++
			} else {
				break
			}
		}

		if ans < tmp {
			ans = tmp
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
