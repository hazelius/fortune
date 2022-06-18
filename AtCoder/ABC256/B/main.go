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
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	p := 0
	bs := make([]int, 4)
	for _, a := range as {
		for i, v := range bs {
			if v > 0 {
				v += a
				if v >= 4 {
					v = 0
					p++
				}
				bs[i] = v
			}
		}
		if a >= 4 {
			p++
		} else {
			for i, v := range bs {
				if v == 0 {
					bs[i] = a
					break
				}
			}
		}
	}
	return p
}

func main() {
	fmt.Println(run(os.Stdin))
}
