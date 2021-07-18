// C - MAD TEAM
// https://atcoder.jp/contests/zone2021/tasks/zone2021_c
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

type member struct {
	a int
	b int
	c int
	d int
	e int
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	members := make([]member, n)
	for i := range members {
		members[i] = member{readInt(), readInt(), readInt(), readInt(), readInt()}
	}

	low, high := 0, int(1e9+1)
	for low+1 < high {
		mid := low + (high-low)/2

		mm := make(map[int]bool)
		for _, one := range members {
			tmp := 0
			if one.a >= mid {
				tmp += 1 << 4
			}
			if one.b >= mid {
				tmp += 1 << 3
			}
			if one.c >= mid {
				tmp += 1 << 2
			}
			if one.d >= mid {
				tmp += 1 << 1
			}
			if one.e >= mid {
				tmp += 1
			}
			mm[tmp] = true
		}
		ret := false

		bitmembers := make([]int, len(mm))
		i := 0
		for m := range mm {
			bitmembers[i] = m
			i++
		}

	b:
		for i := 0; i < len(bitmembers); i++ {
			for j := i; j < len(bitmembers); j++ {
				for k := j; k < len(bitmembers); k++ {
					if bitmembers[i]|bitmembers[j]|bitmembers[k] == 0b11111 {
						ret = true
						break b
					}
				}
			}
		}

		if ret {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	fmt.Println(run(os.Stdin))
}
