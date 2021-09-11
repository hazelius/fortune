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

	alp := "abcdefghijklmnopqrstuvwxyz"
	ans := make([]byte, 26)
	for i := range ans {
		p := readInt()
		ans[i] = alp[p-1]
	}
	return string(ans)
}

func main() {
	fmt.Println(run(os.Stdin))
}
