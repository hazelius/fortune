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

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt() * 2
		ass := make([]int, n)
		for j := range ass {
			ass[j] = readInt()
		}

		amap := make(map[int][2]int, n)
		for j, a := range ass {
			as, ok := amap[a]
			if !ok {
				amap[a] = [2]int{j, 0}
				continue
			}
			if abs(as[0]-j) <= 1 {
				delete(amap, a)
				continue
			}
			as[1] = j
			amap[a] = as
		}

		ans := 0
		for j := 1; j < n; j++ {
			idxs, ok := amap[j]
			if !ok {
				continue
			}
			rmap := make(map[int]bool)
			for _, idx := range idxs {
				if idx > 0 {
					rmap[ass[idx-1]] = true
				}
				if idx < n-1 {
					rmap[ass[idx+1]] = true
				}
			}

			for k := range rmap {
				if k < j {
					continue
				}
				idxs2, ok := amap[k]
				if !ok {
					continue
				}
				if abs(idxs[0]-idxs2[0]) <= 1 {
					if abs(idxs[1]-idxs2[1]) <= 1 {
						ans++
					}
				} else if abs(idxs[0]-idxs2[1]) <= 1 {
					if abs(idxs[1]-idxs2[0]) <= 1 {
						ans++
					}
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	run(os.Stdin, os.Stdout)
}
