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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	smap := make(map[rune][]int, 10)
	for i := 0; i < n; i++ {
		s := readString()
		for j, c := range s {
			idxes, ok := smap[c]
			if !ok {
				idxes = make([]int, 0, n)
			}
			idxes = append(idxes, j)
			smap[c] = idxes
		}
	}

	ans := int(1e9)
	for _, idxes := range smap {
		maxi := 0
		mapi := make(map[int]bool)
		for _, i := range idxes {
			for {
				if mapi[i] {
					i += 10
				} else {
					mapi[i] = true
					break
				}
			}
			if i > maxi {
				maxi = i
			}
		}
		if ans > maxi {
			ans = maxi
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
