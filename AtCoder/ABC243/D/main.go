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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := readInt(), readInt()
	s := readString()
	ido := make([][]int, 0, n)
	idx := -1
	for _, c := range s {
		if c == 'U' {
			if idx < 0 {
				x /= 2
			} else {
				if ido[idx][1] <= 1 {
					if idx < 0 {
						ido = make([][]int, 0, n)
					} else {
						ido = ido[:idx]
					}
					idx--
				} else {
					ido[idx][1]--
				}
			}
		} else {
			rl := 0
			if c == 'L' {
				rl = 1
			}
			if idx < 0 || ido[idx][0] != rl {
				idx++
				ido = append(ido, []int{rl, 1})
			} else {
				ido[idx][1]++
			}
		}
	}

	ans := x
	for _, v := range ido {
		if v[0] < 0 {
			break
		}
		for i := 0; i < v[1]; i++ {
			ans *= 2
			if v[0] == 0 {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
