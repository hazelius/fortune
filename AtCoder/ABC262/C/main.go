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
	amap := make(map[int]bool)
	bmap := make(map[int]int)
	for i := 0; i < n; i++ {
		a := readInt()
		if a == i+1 {
			amap[a] = true
		} else {
			bmap[i+1] = a
		}
	}
	l := len(amap)

	ans := 0
	for k, v := range bmap {
		if b, ok := bmap[v]; ok {
			if b == k {
				ans++
			}
		}
	}
	ans /= 2
	ans += l * (l - 1) / 2

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
