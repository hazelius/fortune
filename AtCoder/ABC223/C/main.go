package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readFloat64() float64 {
	sc.Scan()

	i, _ := strconv.ParseFloat(sc.Text(), 64)
	return i
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	abs := make([][]float64, n)
	var sum float64
	for i := range abs {
		a, b := readFloat64(), readFloat64()
		abs[i] = []float64{a, b}
		sum += a / b
	}

	endtime := sum / 2
	var ans float64
	for _, ab := range abs {
		a, b := ab[0], ab[1]
		l := a / b
		if l >= endtime {
			ans += b * endtime
			break
		}
		endtime -= l
		ans += a
	}
	return fmt.Sprintf("%.15f", ans)
}

func main() {
	fmt.Println(run(os.Stdin))
}
