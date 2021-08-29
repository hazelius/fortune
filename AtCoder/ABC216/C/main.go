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
	ans := ""
	for {
		if n == 0 {
			break
		}
		if n%2 > 0 {
			ans = "A" + ans
			n--
			continue
		}
		ans = "B" + ans
		n /= 2

	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
