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

	_, q := readInt(), readInt()
	hato := make(map[int]int)
	su := make(map[int]int)
	mu := make(map[int]bool)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			p, h := readInt(), readInt()
			no, ok := hato[p]
			if !ok {
				no = p
			}

			cnt, ok := su[no]
			if !ok {
				mu[no] = true
			} else {
				cnt--
				if cnt == 0 {
					delete(su, no)
				} else {
					su[no] = cnt
				}
			}

			cnt, ok = su[h]
			if !ok {
				if mu[h] {
					delete(mu, h)
				} else {
					su[h] = 1
				}
			} else {
				su[h] = cnt + 1
			}
			hato[p] = h
		case 2:
			fmt.Fprintln(out, len(su))
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
