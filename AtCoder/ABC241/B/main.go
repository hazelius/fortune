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

	n, m := readInt(), readInt()
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a]++
	}

	for i := 0; i < m; i++ {
		b := readInt()
		cnt, ok := amap[b]
		if !ok || cnt == 0 {
			return "No"
		}
		amap[b]--
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
