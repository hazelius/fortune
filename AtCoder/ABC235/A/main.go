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

	abc := readInt()
	ans := abc
	abc = (abc*10 + abc/100) % 1000
	ans += abc
	abc = (abc*10 + abc/100) % 1000
	ans += abc
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
