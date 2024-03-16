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

	n, h, w := readInt(), readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		abs[i] = []int{readInt(), readInt()}
	}

	for idx := range abs {
		used := make([]bool, n)
		tiles := make([][]bool, h)
		for i := range tiles {
			tiles[i] = make([]bool, w)
		}
		if f(idx, 0, 0, abs, tiles, used) {
			fmt.Fprint(out, "Yes")
			return
		}
	}

	fmt.Fprint(out, "No")
}

func f(idx, x, y int, abs [][]int, tiles [][]bool, used []bool) bool {
	if used[idx] {
		return false
	}
	used[idx] = true

	ab := abs[idx]
	ar := [][]int{ab}
	if ab[0] != ab[1] {
		ar = append(ar, []int{ab[1], ab[0]})
	}

	for _, ab = range ar {
		if y+ab[1] > len(tiles) {
			continue
		}
		if x+ab[0] > len(tiles[0]) {
			continue
		}
		if tiles[y][x+ab[0]-1] {
			continue
		}

		for i := 0; i < ab[1]; i++ {
			for j := 0; j < ab[0]; j++ {
				tiles[y+i][x+j] = true
			}
		}

		// for _, tt := range tiles {
		// 	for _, v := range tt {
		// 		if v {
		// 			fmt.Print("*")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println("")
		// }
		// fmt.Println("")

		newx, newy := xy(tiles)
		if newx < 0 {
			return true
		}

		for i, v := range used {
			if !v {
				if f(i, newx, newy, abs, tiles, used) {
					return true
				}
			}
		}
		for i := 0; i < ab[1]; i++ {
			for j := 0; j < ab[0]; j++ {
				tiles[y+i][x+j] = false
			}
		}
	}
	used[idx] = false
	return false
}

func xy(tiles [][]bool) (x, y int) {
	for y, row := range tiles {
		for x, val := range row {
			if !val {
				return x, y
			}
		}
	}
	return -1, -1
}

func main() {
	run(os.Stdin, os.Stdout)
}
