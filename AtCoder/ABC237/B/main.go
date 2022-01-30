package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	h, w := readInt(), readInt()
	ans := make([][]int, w)
	for i := range ans {
		ans[i] = make([]int, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[j][i] = readInt()
		}
	}
	ansstr := make([]string, w)
	for i, v := range ans {
		t := fmt.Sprintf("%v", v)
		ansstr[i] = t[1 : len(t)-1]
	}
	return strings.Join(ansstr, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
