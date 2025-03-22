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

	amap := make(map[int]int)
	for i := 0; i < 7; i++ {
		a := readInt()
		amap[a]++
	}
	f3, f2 := false, false
	for _, v := range amap {
		if !f3 && v > 2 {
			f3 = true
			continue
		}
		if v > 1 {
			f2 = true
		}
	}
	if f3 && f2 {
		fmt.Fprint(out, "Yes")
		return
	}
	fmt.Fprint(out, "No")
}

func main() {
	run(os.Stdin, os.Stdout)
}
