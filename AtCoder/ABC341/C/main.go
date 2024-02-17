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

	h, w, n := readInt(), readInt(), readInt()
	t := readString()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}
	ts := make([][]int, n+1)
	ts[0] = []int{0, 0}
	minh, minw := 0, 0
	for i, c := range t {
		pre := ts[i]
		switch c {
		case 'L':
			ts[i+1] = []int{pre[0], pre[1] - 1}
		case 'R':
			ts[i+1] = []int{pre[0], pre[1] + 1}
		case 'U':
			ts[i+1] = []int{pre[0] - 1, pre[1]}
		case 'D':
			ts[i+1] = []int{pre[0] + 1, pre[1]}
		}
		if ts[i+1][0] < minh {
			minh = ts[i+1][0]
		}
		if ts[i+1][1] < minw {
			minw = ts[i+1][1]
		}
	}

	maxh, maxw := 0, 0
	for i := range ts {
		ts[i][0] -= minh
		ts[i][1] -= minw
		if ts[i][0] > maxh {
			maxh = ts[i][0]
		}
		if ts[i][1] > maxw {
			maxw = ts[i][1]
		}
	}

	ans := 0
	for i := 0; i+maxh < h-1; i++ {
		for j := 0; j+maxw < w-1; j++ {
			flg := true
			for _, v := range ts {
				if ss[v[0]+i][v[1]+j] == '#' {
					flg = false
					break
				}
			}
			if flg {
				ans++
			}
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
