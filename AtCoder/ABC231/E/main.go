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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	ans := 0
	b := 0
	for i, a := range as {
		if i >= n-1 {
			ans += x / a
			break
		}
		next := as[i+1]
		amari := x % next
		ret1 := amari / a
		ret2 := b + (next-amari)/a
		b = 0
		if ret1 <= ret2 {
			x -= amari
			ans += ret1
			if ret1 == ret2 {
				b = -1
			}
		} else {
			x += (next - amari)
			ans += ret2
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
