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

type Ints []int

func (a Ints) Len() int           { return len(a) }
func (a Ints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints) Less(i, j int) bool { return abs(a[i]) < abs(a[j]) }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		as := make([]int, n)
		for j := range as {
			as[j] = readInt()
		}
		mapa := make(map[int]int)
		for _, a := range as {
			mapa[a]++
		}
		if len(mapa) == 1 {
			fmt.Fprintln(out, "Yes")
			continue
		}
		if len(mapa) == 2 {
			a1 := as[0]
			flg := true
			for k := range mapa {
				if k != a1 && k != -a1 {
					flg = false
				}
			}
			if flg {
				cnt1 := mapa[a1]
				cnt2 := mapa[-a1]
				if abs(cnt1-cnt2) <= 1 {
					fmt.Fprintln(out, "Yes")
					continue
				}
			}
		}

		sort.Sort(Ints(as))
		flg := true
		for j := 2; j < n; j++ {
			if as[j]*as[j-2] != as[j-1]*as[j-1] {
				flg = false
				break
			}
		}
		if flg {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
