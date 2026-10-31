package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	as := make([]string, 5)
	for i := range as {
		as[i] = readString()
	}
	maxCnt := 0
	cntX := make([]int, 5)
	cnt2 := make([]int, 2)
	for i, v := range as {
		cnt := 0
		for j, vv := range v {
			if vv == 'o' {
				cnt++
				cntX[j]++
			}
		}
		if cnt > maxCnt {
			maxCnt = cnt
		}
		if v[i] == 'o' {
			cnt2[0]++
		}
		if v[4-i] == 'o' {
			cnt2[1]++
		}
	}

	for _, v := range cntX {
		if v > maxCnt {
			maxCnt = v
		}
	}
	for _, v := range cnt2 {
		if v > maxCnt {
			maxCnt = v
		}
	}

	fmt.Fprint(out, 5-maxCnt)
}

func main() {
	run(os.Stdin, os.Stdout)
}
