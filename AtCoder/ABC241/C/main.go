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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = readString()
	}

	tate := make([][]int, n)
	yoko := make([][]int, n)
	naname1 := make([][]int, n)
	naname2 := make([][]int, n)
	for y := 0; y < n; y++ {
		tate[y] = make([]int, n)
		yoko[y] = make([]int, n)
		naname1[y] = make([]int, n)
		naname2[y] = make([]int, n)
		for x := 0; x < n; x++ {
			if x > 0 {
				yoko[y][x] = yoko[y][x-1]
				if x >= 6 && ss[y][x-6] == '#' {
					yoko[y][x]--
				}
			}
			if y > 0 {
				tate[y][x] = tate[y-1][x]
				if y >= 6 && ss[y-6][x] == '#' {
					tate[y][x]--
				}
			}
			if x > 0 && y > 0 {
				naname1[y][x] = naname1[y-1][x-1]
				if x >= 6 && y >= 6 && ss[y-6][x-6] == '#' {
					naname1[y][x]--
				}
			}
			if x < n-1 && y > 0 {
				naname2[y][x] = naname2[y-1][x+1]
				if x < n-6 && y >= 6 && ss[y-6][x+6] == '#' {
					naname2[y][x]--
				}
			}

			if ss[y][x] == '#' {
				tate[y][x]++
				yoko[y][x]++
				naname1[y][x]++
				naname2[y][x]++
			}

			if tate[y][x] >= 4 ||
				yoko[y][x] >= 4 {
				return "Yes"
			}
			if x >= 5 && y >= 5 && naname1[y][x] >= 4 {
				return "Yes"
			}
			if x < n-5 && y >= 5 && naname2[y][x] >= 4 {
				return "Yes"
			}
		}
	}
	return "No"
}

func main() {
	fmt.Println(run(os.Stdin))
}
