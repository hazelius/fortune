package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(n string) string {
	str := strings.Trim(n, "0")
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return "No"
		}
	}

	return "Yes"
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readString()
	fmt.Println(run(n))
}
