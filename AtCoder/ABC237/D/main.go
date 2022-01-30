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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()
	ans := make([]int, n*2+1)

	ans[n] = n
	ridx, lidx := n, n
	for i := n - 1; i > 0; i-- {
		if s[i] == 'L' {
			for ans[lidx] > 0 {
				lidx++
			}
			ans[lidx] = i
		} else {
			for ans[ridx] > 0 {
				ridx--
			}
			ans[ridx] = i
		}
	}

	start := 0
	for i, v := range ans {
		if v > 0 {
			start = i
			break
		}
	}
	end := 0
	for i := start + 1; i < len(ans); i++ {
		if ans[i] == 0 {
			end = i
			break
		}
	}
	if s[0] == 'L' {
		end++
	} else {
		start--
	}

	ansstr := fmt.Sprintf("%v", ans[start:end])
	return fmt.Sprintf("%v", ansstr[1:len(ansstr)-1])
}

func main() {
	fmt.Println(run(os.Stdin))
}
