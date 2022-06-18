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

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	hs := []int{readInt(), readInt(), readInt()}
	ws := []int{readInt(), readInt(), readInt()}

	ans := 0
	for i11 := 1; i11 < hs[0]-1; i11++ {
		for i12 := 1; i12 < hs[0]-i11; i12++ {
			i13 := hs[0] - i11 - i12
			for i21 := 1; i21 < ws[0]-i11; i21++ {
				i31 := ws[0] - i11 - i21
				for i22 := 1; i22 < hs[1]-i21 && i22 < ws[1]-i12; i22++ {
					i23 := hs[1] - i21 - i22
					i32 := ws[1] - i12 - i22
					i33 := hs[2] - i31 - i32
					if i33 > 0 && i31+i32+i33 == hs[2] && i13+i23+i33 == ws[2] {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
