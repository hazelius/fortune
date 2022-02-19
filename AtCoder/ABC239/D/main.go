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

	a, b, c, d := readInt(), readInt(), readInt(), readInt()
	prs := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211}

	for i := a; i <= b; i++ {
		flg := false
		for _, p := range prs {
			s := p - i
			if s < c || d < s {
				continue
			}
			flg = true
		}
		if !flg {
			return "Takahashi"
		}
	}

	return "Aoki"
}

func main() {
	fmt.Println(run(os.Stdin))
}
