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

	n := readInt()
	amap := make(map[int]int)

	ans1, ans2 := 0, 0
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = i
	}
	for i := 0; i < n; i++ {
		b := readInt()
		idx, ok := amap[b]
		if ok {
			if idx == i {
				ans1++
			} else {
				ans2++
			}
		}
	}
	return fmt.Sprintf("%v %v", ans1, ans2)
}

func main() {
	fmt.Println(run(os.Stdin))
}
