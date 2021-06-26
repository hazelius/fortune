package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type lr struct {
	l float64
	r float64
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	list := make([]lr, n)
	for i := 0; i < n; i++ {
		switch readInt() {
		case 1:
			list[i] = lr{readFloat(), readFloat()}
		case 2:
			list[i] = lr{readFloat(), readFloat() - 0.5}
		case 3:
			list[i] = lr{readFloat() + 0.5, readFloat()}
		case 4:
			list[i] = lr{readFloat() + 0.5, readFloat() - 0.5}
		}
	}

	ans := 0
	for i, v := range list {
		for j := i + 1; j < n; j++ {
			if math.Max(v.l, list[j].l) <= math.Min(v.r, list[j].r) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
