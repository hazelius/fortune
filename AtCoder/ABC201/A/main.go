// https://atcoder.jp/contests/abc200/tasks/abc200_a
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(a1, a2, a3 int) string {
	a := []int{a1, a2, a3}
	sort.Ints(a)
	if a[1]-a[0] == a[2]-a[1] {
		return "Yes"
	}
	return "No"
}

func main() {
	sc.Split(bufio.ScanWords)
	a1, a2, a3 := readInt(), readInt(), readInt()
	fmt.Println(run(a1, a2, a3))
}
