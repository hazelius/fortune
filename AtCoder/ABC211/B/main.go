package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	ss := make([]string, 4)
	for i := range ss {
		ss[i] = readString()
	}

	cnt := 0
	hits := []string{"H", "2B", "3B", "HR"}
	for _, h := range hits {
		for _, v := range ss {
			if v == h {
				cnt++
				break
			}
		}
	}

	if cnt < 4 {
		return "No"
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
