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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	cuts := make([]bool, 361)
	a := 0
	cuts[360] = true
	for i := 0; i < n; i++ {
		a = (a + readInt()) % 360
		cuts[a] = true
	}

	ans := 0
	d := 0
	for i, v := range cuts {
		if v {
			m := i - d
			if ans < m {
				ans = m
			}
			d = i
		}
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
