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

	n := readInt()
	sq := make(map[int]bool)
	for i := 1; i <= n; i++ {
		sq[i*i] = true
	}

	d := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			if _, ok := d[j]; !ok {
				d[j] = make(map[int]bool)
			}
			d[j][i] = true
		}
	}

	cnt := make([]int, n+1)
	for i, v := range d {
		f := 0
		for j := range v {
			if sq[j] && f < j {
				f = j
			}
		}
		cnt[i/f]++
	}

	ans := 0
	for _, v := range cnt {
		ans += v * v
	}
	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
