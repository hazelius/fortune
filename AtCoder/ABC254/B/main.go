package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, i+1)
		for j := range ans[i] {
			if j == 0 || j == i {
				ans[i][j] = 1
				continue
			}
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	at := fmt.Sprintf("%v", ans)
	at = strings.ReplaceAll(at, "] [", "\n")
	return at[2 : len(at)-2]
}

func main() {
	fmt.Println(run(os.Stdin))
}
