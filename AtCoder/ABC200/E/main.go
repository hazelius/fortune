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
	for i := 3; i < n; i++ {
		v := i * i * i
		if k < v {
			k -= v
			continue
		}

	}
	return n
}

// (1,1,1)
// (1,1,2) (1,2,1) (2,1,1)
// (1,1,3) (1,2,2) (2,1,1) (1,3,1) (2,2,1) (3,1,1)
// (1,1,4) (1,2,3) (2,1,3) (1,3,2) (2,2,2) (1,4,1) (2,3,1) (3,2,1) (4,1,1)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()
	fmt.Println(run(n, k))
}
