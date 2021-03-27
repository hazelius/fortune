package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s string) string {
	return s[1:] + s[:1]
}

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
