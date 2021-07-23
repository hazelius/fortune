// C - Digital Graffiti
// https://atcoder.jp/contests/abc191/tasks/abc191_c
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

	h, w := readInt(), readInt()
	ss := make([]string, h)
	for i := range ss {
		ss[i] = readString()
	}

	ans := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			cnt := 0
			if ss[i][j] == '#' {
				cnt++
			}
			if ss[i][j+1] == '#' {
				cnt++
			}
			if ss[i+1][j] == '#' {
				cnt++
			}
			if ss[i+1][j+1] == '#' {
				cnt++
			}
			if cnt == 1 || cnt == 3 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
