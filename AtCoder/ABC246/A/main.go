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

	xmap, ymap := make(map[int]int), make(map[int]int)
	for i := 0; i < 3; i++ {
		x, y := readInt(), readInt()
		xmap[x]++
		ymap[y]++
	}
	x, y := 0, 0
	for i, v := range xmap {
		if v == 1 {
			x = i
			break
		}
	}
	for i, v := range ymap {
		if v == 1 {
			y = i
			break
		}
	}
	return fmt.Sprintf("%v %v", x, y)
}

func main() {
	fmt.Println(run(os.Stdin))
}
