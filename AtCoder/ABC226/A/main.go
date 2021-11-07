package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x := readString()
	xs := strings.Split(x, ".")
	ans, _ := strconv.Atoi(xs[0])
	dec, _ := strconv.Atoi(xs[1])
	if dec >= 500 {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
