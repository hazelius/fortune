package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int) string {
	p := (n * 108) / 100
	if p == 206 {
		return "so-so"
	} else if p > 206 {
		return ":("
	}
	return "Yay!"
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := readInt()
	fmt.Println(run(n))
}
