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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, q := readInt(), readInt()
	amap := make(map[int][]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = append(amap[a], i+1)
	}

	ans := make([]int, q)
	for i := range ans {
		x, k := readInt(), readInt()
		as := amap[x]
		if k > len(as) {
			ans[i] = -1
		} else {
			ans[i] = as[k-1]
		}
	}
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
