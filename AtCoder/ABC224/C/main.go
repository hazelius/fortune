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
	xys := make([][]int, n)
	for i := range xys {
		xys[i] = []int{readInt(), readInt()}
	}

	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				s := (xys[j][0]-xys[i][0])*(xys[k][1]-xys[i][1]) - (xys[k][0]-xys[i][0])*(xys[j][1]-xys[i][1])
				if s != 0 {
					ans++
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
