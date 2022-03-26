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

	n, k := readInt(), readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	bs := make([]int, n)
	for i := range bs {
		bs[i] = readInt()
	}

	flga, flgb := true, true
	for i := 0; i < n-1; i++ {
		ba, bb := false, false
		if flga {
			if abs(as[i]-as[i+1]) <= k {
				ba = true
			}
			if abs(as[i]-bs[i+1]) <= k {
				bb = true
			}
		}
		if flgb {
			if abs(bs[i]-as[i+1]) <= k {
				ba = true
			}
			if abs(bs[i]-bs[i+1]) <= k {
				bb = true
			}
		}
		flga = ba
		flgb = bb
		if !flga && !flgb {
			return "No"
		}
	}
	return "Yes"
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func main() {
	fmt.Println(run(os.Stdin))
}
