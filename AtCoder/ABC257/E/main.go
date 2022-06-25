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
	cs := make([]int, 9)
	minc := -1
	mincIdx := 0
	for i := range cs {
		c := readInt()
		cs[i] = c
		if minc < 0 || minc >= c {
			minc = c
			mincIdx = i
		}
	}
	cnt := n / minc
	yosan := n % minc
	ans := make([]int, cnt)
	for idx := 0; idx < cnt; idx++ {
		flg := false
		for i := 8; i > mincIdx; i-- {
			if cs[i] <= minc+yosan {
				ans[idx] = i + 1
				yosan -= cs[i] - minc
				flg = true
				break
			}
		}
		if flg {
			continue
		}
		for i := idx; i < cnt; i++ {
			ans[i] = mincIdx + 1
		}
		break
	}
	ansstr := fmt.Sprintf("%v", ans)
	ansstr = strings.ReplaceAll(ansstr, " ", "")
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
