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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	cards := make(map[int]int)
	for i := 0; i < 5; i++ {
		a := readInt()
		cards[a]++
	}
	if len(cards) != 2 {
		fmt.Fprint(out, "No")
		return
	}
	for _, v := range cards {
		if !(v == 2 || v == 3) {
			fmt.Fprint(out, "No")
			return
		}
	}
	fmt.Fprint(out, "Yes")
}

func main() {
	run(os.Stdin, os.Stdout)
}
