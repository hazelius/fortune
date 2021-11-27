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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(r io.Reader) int {
	sc = bufio.NewScanner(r)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	k := readInt()

	cnt := strings.Count(s, ".")
	if cnt <= k {
		return len(s)
	}

	idxes := make([]int, cnt)

	ans := 0
	idx := -1
	for i, x := range s {
		if x == 'X' {
			continue
		}
		idx++
		idxes[idx] = i

		j := idx - k
		if j < 0 {
			continue
		}

		l := i
		if j > 0 {
			l = i - (idxes[j-1] + 1)
		}
		if ans < l {
			ans = l
		}
	}

	i := cnt - k

	l := len(s)
	if i > 0 {
		l = l - (idxes[i-1] + 1)
	}
	if ans < l {
		ans = l
	}

	return ans
}

func main() {
	fmt.Println(run(os.Stdin))
}
