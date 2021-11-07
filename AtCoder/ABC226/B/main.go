package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	lmap := make(map[string]bool)
	for i := 0; i < n; i++ {
		l := readInt()
		ls := make([]int, l)
		for i := range ls {
			ls[i] = readInt()
		}
		str := fmt.Sprintf("%v", ls)
		lmap[str] = true
	}
	return len(lmap)
}

func main() {
	fmt.Println(run(os.Stdin))
}
