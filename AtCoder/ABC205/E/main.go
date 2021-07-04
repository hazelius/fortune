// E - White and Black Balls
// https://atcoder.jp/contests/abc205/tasks/abc205_e
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var mod int = 1e9 + 7

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m, k := readInt(), readInt(), readInt()

	if n > m+k {
		return 0
	}

	ans := combination(n+m, n)
	if n > k {
		ans -= combination(n+m, m+k+1)
		if ans < 0 {
			ans += mod
		}
	}
	return ans
}

func combination(n, k int) int {
	if n-k < k {
		k = n - k
	}

	v := 1
	for i := 0; i < k; i++ {
		v *= n - i
		v %= mod
		v *= pow(i+1, mod-2)
		v %= mod
	}
	return v
}

func pow(x, y int) int {
	res := 1
	for y > 0 {
		if y&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}

func main() {
	fmt.Println(run(os.Stdin))
}
