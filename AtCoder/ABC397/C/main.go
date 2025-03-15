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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	amap1 := make(map[int]int)
	for _, a := range as {
		amap1[a]++
	}

	ans := len(amap1)

	amap2 := make(map[int]int)
	for _, a := range as {
		cnt := amap1[a]
		cnt--
		amap1[a] = cnt
		if cnt == 0 {
			delete(amap1, a)
		}
		amap2[a]++

		tmp := len(amap1) + len(amap2)
		if ans < tmp {
			ans = tmp
		}
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
