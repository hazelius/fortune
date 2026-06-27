package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	adb := make([][]int, n)
	for i := range adb {
		adb[i] = []int{readInt(), readInt(), readInt()}
	}
	sort.Slice(adb, func(i, j int) bool {
		return adb[i][1] < adb[j][1]
	})

	colorCnt := make(map[int]int)
	for _, v := range adb {
		colorCnt[v[0]]++
	}

	idx := 0
	for i := 0; i < m; i++ {
		for ; idx < n; idx++ {
			if adb[idx][1]-1 != i {
				break
			}
			a := adb[idx][0]
			b := adb[idx][2]
			colorCnt[a]--
			colorCnt[b]++
			if colorCnt[a] == 0 {
				delete(colorCnt, a)
			}
		}
		fmt.Fprintln(out, len(colorCnt))
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
