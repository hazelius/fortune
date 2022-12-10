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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	s := readString()

	regex := regexp.MustCompile(`^[A-Z][0-9]{6}[A-Z]$`)

	if !regex.MatchString(s) {
		fmt.Fprint(out, "No")
		return
	}
	n, err := strconv.Atoi(s[1 : len(s)-1])
	if err != nil {
		fmt.Fprint(out, "No")
		return
	}
	if 100000 <= n && n <= 999999 {
		fmt.Fprint(out, "Yes")
		return
	}
	fmt.Fprint(out, "No")

}

func main() {
	run(os.Stdin, os.Stdout)
}
