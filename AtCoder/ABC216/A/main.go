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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readString()
	nums := strings.Split(n, ".")
	s, _ := strconv.Atoi(nums[1])
	ans := nums[0]
	if s >= 7 {
		ans += "+"
	} else if s < 3 {
		ans += "-"
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
