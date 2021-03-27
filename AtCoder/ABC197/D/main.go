package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readFloat64() float64 {
	sc.Scan()

	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func run(n, x0, y0, xn, yn float64) string {

	var xc float64 = (x0 - xn) / 2
	var yc float64 = (y0 - yn) / 2
	var rad float64 = 2 * math.Pi / n

	var ansx float64 = xc*math.Cos(rad) - yc*math.Sin(rad) + (x0+xn)/2
	var ansy float64 = xc*math.Sin(rad) + yc*math.Cos(rad) + (y0+yn)/2

	return fmt.Sprintf("%.7f %.7f", ansx, ansy)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readFloat64()
	x0, y0 := readFloat64(), readFloat64()
	xn, yn := readFloat64(), readFloat64()

	fmt.Println(run(n, x0, y0, xn, yn))
}
