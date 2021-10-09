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

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

type Person struct {
	no    int
	gcp   string
	point int
}

type Persons []Person

func (a Persons) Len() int      { return len(a) }
func (a Persons) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Persons) Less(i, j int) bool {
	if a[i].point == a[j].point {
		return a[i].no < a[j].no
	}
	return a[i].point > a[j].point
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	ps := make([]Person, n*2)
	for i := range ps {
		ps[i].no = i
		ps[i].gcp = readString()
	}

	for round := 0; round < m; round++ {
		for i := 0; i < n; i++ {
			idx := i * 2
			result := janken(ps[idx].gcp[round], ps[idx+1].gcp[round])
			if result > 0 {
				ps[idx].point++
			} else if result < 0 {
				ps[idx+1].point++
			}
		}
		sort.Sort(Persons(ps))
		// fmt.Println(ps)
	}

	ans := make([]int, len(ps))
	for i, p := range ps {
		ans[i] = p.no + 1
	}
	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func janken(a1, a2 byte) int {
	if a1 == a2 {
		return 0
	}
	if a1 == byte('G') && a2 == byte('C') || a1 == byte('C') && a2 == byte('P') || a1 == byte('P') && a2 == byte('G') {
		return 1
	}
	return -1
}

func main() {
	fmt.Println(run(os.Stdin))
}
