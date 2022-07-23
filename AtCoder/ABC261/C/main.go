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
	sc.Split(bufio.ScanWords)

	n := readInt()
	ans := make([]string, n)
	smap := make(map[string]int)
	for i := range ans {
		s := readString()
		cnt, ok := smap[s]
		if !ok {
			ans[i] = s
			smap[s] = 0
		} else {
			cnt++
			ans[i] = fmt.Sprintf("%v(%v)", s, cnt)
			smap[s] = cnt
		}
	}

	ansstr := fmt.Sprintf("%v", ans)
	return ansstr[1 : len(ansstr)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
