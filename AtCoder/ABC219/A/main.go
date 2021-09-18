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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	x := readInt()
	if 40 > x {
		return fmt.Sprint(40 - x)
	}
	if 70 > x {
		return fmt.Sprint(70 - x)
	}
	if 90 > x {
		return fmt.Sprint(90 - x)
	}

	return "expert"
}

func main() {
	fmt.Println(run(os.Stdin))
}
