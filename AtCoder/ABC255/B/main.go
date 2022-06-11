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

	n, k := readInt(), readInt()
	amap := make(map[int]bool)
	for i := 0; i < k; i++ {
		a := readInt()
		amap[a-1] = true
	}

	akari := make([][]int, k)
	ku := make([][]int, n-k)
	idx1, idx2 := 0, 0
	for i := 0; i < n; i++ {
		xy := []int{readInt(), readInt()}
		if amap[i] {
			akari[idx1] = xy
			idx1++
		} else {
			ku[idx2] = xy
			idx2++
		}
	}

	ans := 0
	for _, xy := range ku {
		mn := -1
		for _, xy2 := range akari {
			x := xy[0] - xy2[0]
			y := xy[1] - xy2[1]
			d := (x * x) + (y * y)
			if mn < 0 || mn > d {
				mn = d
			}
		}
		if mn > ans {
			ans = mn
		}
	}

	return math.Sqrt(float64(ans))
}

func main() {
	fmt.Println(run(os.Stdin))
}
