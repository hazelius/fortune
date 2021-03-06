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

func run(n int) string {
	var r float64
	for i := 1; i < n; i++ {
		r += kitai(i, n)
	}
	return fmt.Sprint(r)
}

func kitai(i, n int) float64 {
	return float64(n) / float64(n-i)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	fmt.Println(run(n))
}
