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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	k := readInt()
	a, b := readInt(), readInt()

	adec, bdec := 0, 0
	ktemp := 1
	for i := 0; a > 0 || b > 0; i++ {
		ca := a % 10
		cb := b % 10

		adec += ktemp * ca
		bdec += ktemp * cb

		ktemp *= k
		a /= 10
		b /= 10
	}

	return adec * bdec
}

func main() {
	fmt.Println(run(os.Stdin))
}
