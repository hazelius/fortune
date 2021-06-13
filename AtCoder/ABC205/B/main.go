package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) string {
	sort.Ints(as)
	for i, a := range as {
		if a != i+1 {
			return "No"
		}
	}
	return "Yes"
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
