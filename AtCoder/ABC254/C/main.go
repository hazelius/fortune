package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	ass := make([][]int, k)
	amari := n % k
	for i := range ass {
		lng := n / k
		if amari > 0 {
			lng++
			amari--
		}
		ass[i] = make([]int, lng)
	}

	for i := 0; i < n; i++ {
		ass[i%k][i/k] = readInt()
	}

	for i := range ass {
		sort.Ints(ass[i])
	}

	for i := 0; i < n-1; i++ {
		nx := i + 1
		if ass[i%k][i/k] > ass[nx%k][nx/k] {
			return "No"
		}
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
