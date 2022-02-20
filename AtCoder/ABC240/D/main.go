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

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	arenzoku := make([][]int, n)
	prea := 0
	idx := -1

	ans := make([]int, n)
	sum := 0

	for i, a := range as {
		if prea != a {
			prea = a
			idx++
			arenzoku[idx] = []int{a, 1}
		} else {
			arenzoku[idx][1]++
			if arenzoku[idx][1] == a {
				sum -= a
				arenzoku[idx] = []int{0, 0}
				idx--
				prea = 0
				if idx >= 0 {
					prea = arenzoku[idx][0]
				}
			}
		}
		sum++
		ans[i] = sum
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
