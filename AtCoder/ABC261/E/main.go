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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, c := readInt(), readInt()
	tas := make([][]int, n)
	for i := range tas {
		tas[i] = []int{readInt(), readInt()}
	}

	ans := make([]int, n)
	for i := 0; i < 30; i++ {
		cb := (c >> i) & 1
		f := []int{0, 1}
		for j, ta := range tas {
			t, a := ta[0], ta[1]
			ab := (a >> i) & 1
			switch t {
			case 1:
				f = []int{f[0] & ab, f[1] & ab}
			case 2:
				f = []int{f[0] | ab, f[1] | ab}
			case 3:
				f = []int{f[0] ^ ab, f[1] ^ ab}
			}
			cb = f[cb]
			ans[j] |= cb << i
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
