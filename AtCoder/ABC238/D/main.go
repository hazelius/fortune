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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	t := readInt()
	anss := make([]string, t)
	for i := 0; i < t; i++ {
		a, s := readInt(), readInt()
		nokori := s - (a * 2)
		if nokori < 0 || nokori&a > 0 {
			anss[i] = "No"
		} else {
			anss[i] = "Yes"
		}
	}

	return strings.Join(anss, " ")
}

func main() {
	fmt.Println(run(os.Stdin))
}
