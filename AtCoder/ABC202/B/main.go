package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s string) string {
	var ans string
	for _, v := range s {
		switch v {
		case '6':
			ans = "9" + ans
		case '9':
			ans = "6" + ans
		default:
			ans = string(v) + ans
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	s := readString()
	fmt.Println(run(s))
}
