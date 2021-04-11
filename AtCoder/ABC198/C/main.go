package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r, x, y int) float64 {
	d := math.Sqrt(float64(x*x + y*y))
	ans := d / float64(r)
	if ans < 1 {
		ans += 1
	}

	return math.Ceil(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	r, x, y := readInt(), readInt(), readInt()
	fmt.Println(run(r, x, y))
}
