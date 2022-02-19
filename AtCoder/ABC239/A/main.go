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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) float64 {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h := readInt()
	ans := h * (h + 12800000)

	return math.Sqrt(float64(ans))
}

func main() {
	fmt.Println(run(os.Stdin))
}
