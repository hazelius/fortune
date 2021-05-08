package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n, k int) int {
	for n%200 == 0 {
		n = n / 200
		k--
	}
	for i := 1; i < k; i += 2 {
		n = n*5 + 1
	}

	if k%2 > 0 {
		n = n*1000 + 200
	}

	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	fmt.Println(run(n, k))
}
