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

	n, m := readInt(), readInt()
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i + 1
	}
	s := fmt.Sprintf("%v", ans)
	fmt.Fprintln(out, s[1:len(s)-1])
	for {
		v := ans[n-1]
		if v < m {
			ans[n-1]++
		} else {
			if n == 1 {
				return
			}
			for i := n - 2; i >= 0; i-- {
				if ans[i] > m-(n-i) {
					if i == 0 {
						return
					}
					continue
				}
				ans[i]++
				for j := i + 1; j < n; j++ {
					ans[j] = ans[j-1] + 1
				}
				break
			}
		}

		s := fmt.Sprintf("%v", ans)
		fmt.Fprintln(out, s[1:len(s)-1])
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
