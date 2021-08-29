package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	names := make([]string, n)
	for i := range names {
		names[i] = readString() + " " + readString()
	}
	sort.Strings(names)

	p := ""
	for _, v := range names {
		if p == v {
			return "Yes"
		}
		p = v
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
