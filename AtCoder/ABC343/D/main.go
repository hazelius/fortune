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

	n, t := readInt(), readInt()
	as := make([]int, n)
	amap := make(map[int]int)
	amap[0] = n
	for i := 0; i < t; i++ {
		a, b := readInt()-1, readInt()
		point := as[a]
		cnt := amap[point]
		cnt--
		if cnt == 0 {
			delete(amap, point)
		} else {
			amap[point] = cnt
		}

		point += b
		as[a] = point
		cnt = amap[point]
		amap[point] = cnt + 1

		fmt.Fprintln(out, len(amap))
	}

}

func main() {
	run(os.Stdin, os.Stdout)
}
