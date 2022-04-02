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

func readFloat() float64 {
	sc.Scan()
	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	a, b := readFloat(), readFloat()
	ra := a / math.Sqrt(a*a+b*b)
	rb := b / math.Sqrt(a*a+b*b)

	return fmt.Sprintf("%.12f %.12f", ra, rb)
}

func main() {
	fmt.Println(run(os.Stdin))
}
