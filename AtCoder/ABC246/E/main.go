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

	n := readInt()
	ax, ay := readInt()-1, readInt()-1
	bx, by := readInt()-1, readInt()-1

	xys := make([][]int, n)
	for i := range xys {
		xys[i] = make([]int, n)
		s := readString()
		for j, c := range s {
			if c == '#' {
				xys[i][j] = -1
			}
		}
	}
	xys[ax][ay] = -1

	queue := [][]int{{ax, ay}}
	for i := 1; len(queue) > 0; i++ {
		queue = f(i, queue, xys)
		if xys[bx][by] > 0 {
			// for _, xy := range xys {
			// 	for _, v := range xy {
			// 		fmt.Printf("%2d", v)
			// 	}
			// 	fmt.Println(" ")
			// }
			// fmt.Println("----")

			return i
		}
	}
	return -1
}

func f(cnt int, queue [][]int, xys [][]int) [][]int {
	ret := make([][]int, 0)
	dir := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, xy := range queue {
		for _, d := range dir {
			x, y := xy[0], xy[1]
			for {
				x += d[0]
				y += d[1]
				if x < 0 || y < 0 || x >= len(xys) || y >= len(xys) {
					break
				}
				p := xys[x][y]
				if p == 0 {
					xys[x][y] = cnt
					ret = append(ret, []int{x, y})
				} else if p < 0 || p < cnt {
					break
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(run(os.Stdin))
}
