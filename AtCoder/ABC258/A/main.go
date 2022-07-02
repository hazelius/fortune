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

	k := readInt()
	h := 21 + k/60
	s := k % 60

	return fmt.Sprintf("%02v:%02v", h, s)
}

func main() {
	fmt.Println(run(os.Stdin))
}
