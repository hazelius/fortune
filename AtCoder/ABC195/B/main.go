package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(a, b, w int) string {
	w *= 1000
	if w < a {
		return "UNSATISFIABLE"
	}

	min := w / a
	amari := w % a
	if amari != 0 {
		if float64(amari)/float64(min) > float64(b-a) {
			return "UNSATISFIABLE"
		}
	}

	max := w / b
	amari = w % b
	if amari != 0 {
		max++
	}

	return fmt.Sprintf("%v %v", max, min)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, w := readInt(), readInt(), readInt()

	fmt.Println(run(a, b, w))
}
