package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var sc *bufio.Scanner

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()

	regex := regexp.MustCompile(`^[a-z]+$`)
	if regex.MatchString(s) {
		return "No"
	}
	regex = regexp.MustCompile(`^[A-Z]+$`)
	if regex.MatchString(s) {
		return "No"
	}

	maps := make(map[rune]bool)
	for _, c := range s {
		if maps[c] {
			return "No"
		}
		maps[c] = true
	}
	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
