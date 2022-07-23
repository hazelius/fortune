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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([][]string, n)
	for i := range as {
		as[i] = make([]string, n)
		a := readString()
		for j := range as[i] {
			as[i][j] = a[j : j+1]
		}
	}
	for i := range as {
		for j := range as[i] {
			switch as[i][j] {
			case "-":
				continue
			case "W":
				if as[j][i] != "L" {
					return "incorrect"
				}
			case "L":
				if as[j][i] != "W" {
					return "incorrect"
				}
			case "D":
				if as[j][i] != "D" {
					return "incorrect"
				}
			}
		}
	}

	return "correct"
}

func main() {
	fmt.Println(run(os.Stdin))
}
