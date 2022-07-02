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

	n := readInt()
	aas := make([][]int, n)
	maxa := 0
	for i := range aas {
		aas[i] = make([]int, n)
		aa := readInt()
		for j := n - 1; j >= 0; j-- {
			a := aa % 10
			aa /= 10

			if maxa < a {
				maxa = a
			}
			aas[i][j] = a
		}
	}

	ans := 0
	for i, as := range aas {
		for j, a := range as {
			if a < maxa {
				continue
			}
			m := f(i, j, aas)
			if m > ans {
				ans = m
			}
		}
	}

	return ans
}

func f(i, j int, aas [][]int) int {
	n := len(aas)
	dir := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	ret := 0
	for _, v := range dir {
		x, y := i, j
		z := aas[x][y]
		for k := 0; k < n-1; k++ {
			x = (x + v[0]) % n
			if x < 0 {
				x = n - 1
			}
			y = (y + v[1]) % n
			if y < 0 {
				y = n - 1
			}
			z = 10*z + aas[x][y]
		}
		if ret < z {
			ret = z
		}
	}
	return ret
}

func main() {
	fmt.Println(run(os.Stdin))
}
