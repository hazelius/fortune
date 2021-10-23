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

	h, w := readInt(), readInt()
	as := make([][]int, h)
	for i := range as {
		as[i] = make([]int, w)
		for j := range as[i] {
			as[i][j] = readInt()
		}
	}

	for h1 := 0; h1 < h-1; h1++ {
		for w1 := 0; w1 < w-1; w1++ {
			for h2 := h1 + 1; h2 < h; h2++ {
				for w2 := w1 + 1; w2 < w; w2++ {
					if as[h1][w1]+as[h2][w2] > as[h2][w1]+as[h1][w2] {
						return "No"
					}
				}
			}

		}
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
