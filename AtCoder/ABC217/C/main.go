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
	qs := make([]int, n)
	for i := range qs {
		p := readInt() - 1
		qs[p] = i + 1
	}
	ans := fmt.Sprintf("%v", qs)
	return ans[1 : len(ans)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
