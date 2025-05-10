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

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	ans := make([][]byte, h)
	for i := range ss {
		ans[i] = []byte(ss[i])
	}
	scnt := make([][]int, h)
	for i := range scnt {
		scnt[i] = make([]int, w)
	}
	es := make([][]int, 0)
	for i, s := range ss {
		for j, e := range s {
			if e == 'E' {
				es = append(es, []int{i, j})
			}
		}
	}
	f(es, ss, scnt, ans)
	for _, s := range ans {
		fmt.Fprintln(out, string(s))
	}
}

func f(es [][]int, ss []string, scnt [][]int, ans [][]byte) {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ch := []byte{'v', '^', '>', '<'}
	que := make([][]int, 4*len(es))
	idx := 0
	for _, e := range es {
		for j, d := range dir {
			que[idx] = []int{e[0] + d[0], e[1] + d[1], 1, j}
			idx++
		}
	}
	for len(que) > 0 {
		i, j, cnt, dc := que[0][0], que[0][1], que[0][2], que[0][3]
		que = que[1:]
		if i < 0 || i >= len(ss) || j < 0 || j >= len(ss[0]) {
			continue
		}
		if ss[i][j] == 'E' {
			continue
		}
		if ss[i][j] == '#' {
			continue
		}
		if scnt[i][j] == 0 || scnt[i][j] > cnt {
			scnt[i][j] = cnt
			ans[i][j] = ch[dc]
			for idx, d := range dir {
				que = append(que, []int{i + d[0], j + d[1], cnt + 1, idx})
			}
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
