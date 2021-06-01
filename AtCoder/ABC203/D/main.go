// D - Pond
// https://atcoder.jp/contests/abc203/tasks/abc203_d
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

type Square struct {
	n      int
	k      int
	sq     [][]int
	center int
}

func (s *Square) isOK(m int) bool {
	sq := make([][]int, s.n+1)
	sq[0] = make([]int, s.n+1)
	for i, row := range s.sq {
		sq[i+1] = make([]int, s.n+1)
		for j, a := range row {
			if a >= m {
				sq[i+1][j+1] = 1
			}
		}
	}

	for i := 0; i <= s.n; i++ {
		for j := 1; j <= s.n; j++ {
			sq[i][j] += sq[i][j-1]
		}
	}
	for i := 1; i <= s.n; i++ {
		for j := 0; j <= s.n; j++ {
			sq[i][j] += sq[i-1][j]
		}
	}

	for i := 0; i <= s.n-s.k; i++ {
		for j := 0; j <= s.n-s.k; j++ {
			c := sq[i+s.k][j+s.k] - sq[i+s.k][j] - sq[i][j+s.k] + sq[i][j]
			if c < s.center {
				return true
			}
		}
	}

	return false
}

func run(n, k int, sq [][]int) int {
	center := (k * k / 2) + 1
	str := Square{n: n, k: k, sq: sq, center: center}

	return binarySearch(-1, math.MaxInt64-1, &str)
}

type iBinarySearch interface {
	isOK(int) bool
}

func binarySearch(low, high int, o iBinarySearch) int {
	for low+1 < high {
		mid := low + (high-low)/2
		if o.isOK(mid) {
			high = mid
		} else {
			low = mid
		}
	}

	return low
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	sq := make([][]int, n)
	for i := range sq {
		sq[i] = make([]int, n)
		for j := range sq[i] {
			sq[i][j] = readInt()
		}
	}
	fmt.Println(run(n, k, sq))
}
