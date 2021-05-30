package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int) int {
	ans := f(k) * n
	for i := 1; i <= n; i++ {
		ans += 100 * i * k
	}
	return ans
}

func f(k int) int {
	var a int
	for i := 1; i <= k; i++ {
		a += i
	}
	return a
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	fmt.Println(run(n, k))
}
