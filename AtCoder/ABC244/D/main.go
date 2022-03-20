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

	s := []string{readString(), readString(), readString()}
	t := []string{readString(), readString(), readString()}
	cnt := 0
	for i := 0; i < 3; i++ {
		if s[i] == t[i] {
			cnt++
		}
	}
	if cnt == 1 {
		return "No"
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
