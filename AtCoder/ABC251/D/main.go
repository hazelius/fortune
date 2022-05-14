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

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	readInt()

	ans := make([]int, 297)
	for i := 1; i < 100; i++ {
		ans[i-1] = i
		ans[i+98] = i * 100
		ans[i+197] = i * 10000
	}

	ansstr := fmt.Sprintf("%v", ans)
	return fmt.Sprintf("%v\n%v", len(ans), ansstr[1:len(ansstr)-1])
}

func main() {
	fmt.Println(run(os.Stdin))
}
