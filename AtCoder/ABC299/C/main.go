package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	readInt()
	s := readString()

	ans := 0

	re := regexp.MustCompile("-o+")
	matches := re.FindAllString(s, -1)
	for _, match := range matches {
		if ans < len(match) {
			ans = len(match)
		}
	}

	re = regexp.MustCompile("o+-")
	matches = re.FindAllString(s, -1)
	for _, match := range matches {
		if ans < len(match) {
			ans = len(match)
		}
	}

	fmt.Fprint(out, ans-1)
}

func main() {
	run(os.Stdin, os.Stdout)
}
