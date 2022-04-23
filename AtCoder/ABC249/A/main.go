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

	a, b, c, d, e, f, x := readInt(), readInt(), readInt(), readInt(), readInt(), readInt(), readInt()

	ad := (x / (a + c)) * a
	if x%(a+c) < a {
		ad += x % (a + c)
	} else {
		ad += a
	}

	td := (x / (d + f)) * d
	if x%(d+f) < d {
		td += x % (d + f)
	} else {
		td += d
	}

	ansa := ad * b
	anst := td * e

	if ansa == anst {
		return "Draw"
	}
	if ansa > anst {
		return "Takahashi"
	}
	return "Aoki"
}

func main() {
	fmt.Println(run(os.Stdin))
}
