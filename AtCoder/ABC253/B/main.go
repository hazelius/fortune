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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	h, _ := readInt(), readInt()
	koma := make([][]int, 2)
	idx := 0
	for i := 0; i < h; i++ {
		s := readString()
		for j, c := range s {
			if c == 'o' {
				koma[idx] = []int{i, j}
				idx++
			}
		}
	}
	ans := abs(koma[0][0]-koma[1][0]) + abs(koma[0][1]-koma[1][1])
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	fmt.Println(run(os.Stdin))
}
