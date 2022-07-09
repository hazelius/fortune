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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	a, b, d := readInt(), readInt(), readInt()
	rd := float64(d) * (math.Pi / 180)
	xn := float64(a)*math.Cos(rd) - float64(b)*math.Sin(rd)
	yn := float64(a)*math.Sin(rd) + float64(b)*math.Cos(rd)
	return fmt.Sprintf("%v %v", xn, yn)
}

func main() {
	fmt.Println(run(os.Stdin))
}
