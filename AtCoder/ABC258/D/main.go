package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, x := readInt(), readInt()
	base := 0
	ans := -1
	for i := 0; i < n; i++ {
		a, b := readInt(), readInt()
		base += a
		time := base + b*(x-i)
		if ans < 0 || time < ans {
			if x > i {
				ans = time
			}
		}
		base += b
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
