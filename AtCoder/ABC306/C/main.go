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
	mapa := make(map[int]int)
	idx := 0
	for i := 0; i < n*3; i++ {
		a := readInt()
		cnt := mapa[a]
		if cnt == 1 {
			as[idx] = a
			idx++
		}
		cnt++
		mapa[a] = cnt
	}

	ans := fmt.Sprintf("%v", as)
	fmt.Fprint(out, ans[1:len(ans)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
