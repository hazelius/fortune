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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	maxa := 0
	mapIdx := make(map[int]bool)
	for i := 0; i < n; i++ {
		a := readInt()
		if maxa < a {
			maxa = a
			mapIdx = make(map[int]bool)
			mapIdx[i+1] = true
		} else if maxa == a {
			mapIdx[i+1] = true
		}
	}

	for i := 0; i < k; i++ {
		b := readInt()
		if mapIdx[b] {
			return "Yes"
		}
	}

	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
