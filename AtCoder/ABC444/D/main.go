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
	sort.Ints(as)

	idx := 0
	ans := ""
	kr := 0
	cnt := n
	for i := 1; ; i++ {
		for as[idx] < i {
			cnt--
			idx++
			if idx >= n {
				break
			}
		}
		if cnt == 0 {
			if kr > 0 {
				ans = strconv.Itoa(kr) + ans
			}
			break
		}
		ans = strconv.Itoa((kr+cnt)%10) + ans
		kr = (kr + cnt) / 10
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
