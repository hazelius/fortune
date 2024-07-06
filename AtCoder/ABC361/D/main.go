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

type query struct {
	s   string
	idx int
	cnt int
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s, t := readString(), readString()
	stmap := make(map[string]bool)

	que := []query{{s + "..", n, 0}}
	for len(que) > 0 {
		q := que[0]
		que = que[1:]
		if q.idx == n && q.s[:n] == t {
			fmt.Fprint(out, q.cnt)
			return
		}

		cnt := q.cnt + 1
		for i := 0; i <= n; i++ {
			if i == q.idx || i+1 == q.idx {
				continue
			}
			ns := []byte(q.s)
			ns[q.idx], ns[i] = ns[i], ns[q.idx]
			ns[q.idx+1], ns[i+1] = ns[i+1], ns[q.idx+1]
			nexts := string(ns)
			if stmap[nexts] {
				continue
			}
			stmap[nexts] = true
			que = append(que, query{nexts, i, cnt})
		}
	}

	fmt.Fprint(out, -1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
