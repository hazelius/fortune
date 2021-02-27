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

func run(r [][]int64) string {
	var result int64 = 1 << 32

	for _, rec := range r {
		a := rec[0]
		p := rec[1]
		x := rec[2]

		if x <= a {
			continue
		}
		if result > p {
			result = p
		}
	}
	if result == 1<<32 {
		result = -1
	}

	return fmt.Sprintf("%v", result)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()

	var result int = 1 << 32

	for i := 0; i < n; i++ {
		a, p, x := readInt(), readInt(), readInt()
		if x <= a {
			continue
		}
		if result > p {
			result = p
		}
	}
	if result == 1<<32 {
		result = -1
	}

	fmt.Println(result)
}
