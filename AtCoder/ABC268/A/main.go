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

	amap := make(map[int]bool)
	for i := 0; i < 5; i++ {
		a := readInt()
		amap[a] = true
	}

	fmt.Fprint(out, len(amap))
}

func main() {
	run(os.Stdin, os.Stdout)
}
