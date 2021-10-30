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

	// for i := 0; i < 20; i++ {
	// 	str := ""
	// 	for j := 1; j <= 7; j++ {
	// 		v := i*7 + j
	// 		str = fmt.Sprintf("%v %v", str, v)
	// 	}
	// 	fmt.Println(str)
	// }

	n, m := readInt(), readInt()
	bs := make([][]int, n)
	for i := range bs {
		bs[i] = make([]int, m)
		for j := range bs[i] {
			b := readInt()
			if b == 0 {
				return "No"
			}
			bs[i][j] = b
		}
	}

	for i, row := range bs {
		if i > 0 {
			pre := bs[i-1][0]
			if row[0] != pre+7 {
				return "No"
			}
		}
		for j := 1; j < len(row); j++ {
			if row[j-1]+1 != row[j] {
				return "No"
			}
			if row[j-1]%7 == 0 {
				return "No"
			}
		}
	}

	return "Yes"
}

func main() {
	fmt.Println(run(os.Stdin))
}
