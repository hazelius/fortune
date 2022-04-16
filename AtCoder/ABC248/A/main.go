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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	ans := "0123456789"
	s := readString()
	for _, c := range s {
		ans = strings.Replace(ans, string(c), "", -1)
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
