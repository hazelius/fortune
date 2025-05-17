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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	ans := 0
	cnt := [2]int{0, 0}
	flg := 0
	for i, a := range as {
		if i == 0 {
			continue
		}
		if a > as[i-1] {
			cnt[flg]++
		} else {
			if cnt[0] > 0 && cnt[1] > 0 {
				ans += cnt[0] * cnt[1]
				cnt[1-flg] = 0
			}
			if i > 1 && as[i-1] > as[i-2] {
				flg = 1 - flg
			}
		}
	}
	if cnt[0] > 0 && cnt[1] > 0 {
		ans += cnt[0] * cnt[1]
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
