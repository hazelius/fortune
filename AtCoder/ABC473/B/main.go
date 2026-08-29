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
	amap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		if _, ok := amap[a]; !ok {
			amap[a] = a
		} else {
			delete(amap, a)
		}
	}
	ans := 0
	for _, v := range amap {
		ans += v
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
