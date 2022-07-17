package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

type Math [][]int

func (a Math) Len() int      { return len(a) }
func (a Math) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Math) Less(i, j int) bool {
	if a[i][1] == a[j][1] {
		return a[i][0] < a[j][0]
	}
	return a[i][1] > a[j][1]
}

type Eng [][]int

func (a Eng) Len() int      { return len(a) }
func (a Eng) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Eng) Less(i, j int) bool {
	if a[i][2] == a[j][2] {
		return a[i][0] < a[j][0]
	}
	return a[i][2] > a[j][2]
}

type Sum [][]int

func (a Sum) Len() int      { return len(a) }
func (a Sum) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Sum) Less(i, j int) bool {
	if a[i][3] == a[j][3] {
		return a[i][0] < a[j][0]
	}
	return a[i][3] > a[j][3]
}

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, x, y, z := readInt(), readInt(), readInt(), readInt()
	abs := make([][]int, n)
	for i := range abs {
		a := readInt()
		abs[i] = []int{i + 1, a, 0, 0}
	}

	for i := range abs {
		b := readInt()
		abs[i][2] = b
		abs[i][3] = abs[i][1] + b
	}

	ans := make([]int, x+y+z)

	sort.Sort(Math(abs))
	for i := 0; i < x; i++ {
		ans[i] = abs[i][0]
	}

	abs = abs[x:]
	sort.Sort(Eng(abs))
	for i := 0; i < y; i++ {
		ans[i+x] = abs[i][0]
	}

	abs = abs[y:]
	sort.Sort(Sum(abs))
	for i := 0; i < z; i++ {
		ans[i+x+y] = abs[i][0]
	}

	sort.Ints(ans)
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
