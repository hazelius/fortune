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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	a, b, c := readInt(), readInt(), readInt()
	ans := []int{a, b, c}
	sort.Ints(ans)
	return ans[1] + ans[2]
}

func main() {
	fmt.Println(run(os.Stdin))
}
