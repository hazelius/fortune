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

func run(x string) string {
	ret := strings.Split(x, ".")
	return ret[0]
}

func main() {
	sc.Split(bufio.ScanWords)
	x := readString()

	fmt.Println(run(x))
}
