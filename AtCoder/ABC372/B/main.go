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

	m := readInt()
	ans := make([]int, 0, 20)
	for m > 0 {
		cnt := 0
		tmp := 1
		for tmp <= m {
			tmp *= 3
			cnt++
		}
		m -= tmp / 3
		ans = append(ans, cnt-1)
	}
	str := fmt.Sprintf("%v", ans)
	fmt.Fprintln(out, len(ans))
	fmt.Fprint(out, str[1:len(str)-1])
}

func main() {
	run(os.Stdin, os.Stdout)
}
