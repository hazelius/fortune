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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	ans := make([]int, 0)

	flg := true
	idx := n - 1
	for i := 1; i <= idx; i++ {
		if as[0] == as[i] {
			continue
		}
		if i < idx && as[0] == as[i]+as[idx] {
			idx--
			continue
		}
		flg = false
		break
	}
	if flg {
		ans = append(ans, as[0])
	}

	if n%2 == 0 {
		flg = true
		l := as[0] + as[n-1]
		for i := 1; i < n/2; i++ {
			if as[i]+as[n-i-1] == l {
				continue
			}
			flg = false
			break
		}
		if flg {
			ans = append(ans, l)
		}
	}

	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
