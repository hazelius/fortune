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

	n, x := readInt(), readInt()
	amap := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		amap[a] = true
	}

	for a := range amap {
		if amap[a+x] || amap[a-x] {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
