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

	n, k, q := readInt(), readInt(), readInt()
	as := make([]int, k)
	for i := range as {
		as[i] = readInt() - 1
	}
	ns := make([]bool, n)
	for _, a := range as {
		ns[a] = true
	}

	for i := 0; i < q; i++ {
		l := readInt() - 1
		a := as[l]
		if a < n-1 && !ns[a+1] {
			as[l]++
			ns[a] = false
			ns[a+1] = true
		}
	}

	for i := range as {
		as[i]++
	}
	ans := fmt.Sprintf("%v", as)
	return ans[1 : len(ans)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
