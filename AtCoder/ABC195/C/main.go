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

func run(n int) string {
	result := 0
	for e := 3; e <= 15; e += 3 {
		t := pow(10, e)
		if n < t {
			break
		}

		m := pow(10, e+3) - 1
		if n < m {
			m = n
		}
		c := e / 3

		result += (m - t + 1) * c
	}

	return fmt.Sprint(result)
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()

	fmt.Println(run(n))
}
