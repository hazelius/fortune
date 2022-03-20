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

	readInt()
	t := readString()

	directions := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	d := 0
	ans := []int{0, 0}
	for _, c := range t {
		if c == 'S' {
			ans = []int{ans[0] + directions[d][0], ans[1] + directions[d][1]}
		} else {
			d = (d + 1) % 4
		}
	}

	return fmt.Sprintf("%v %v", ans[0], ans[1])
}

func main() {
	fmt.Println(run(os.Stdin))
}
