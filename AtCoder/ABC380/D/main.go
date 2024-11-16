package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
	"strconv"
	"unicode"
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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := readString()
	sre := reverseCase(s)
	q := readInt()
	ans := make([]string, q)
	leng := len(s)
	for i := range ans {
		k := readInt() - 1
		idx := k % leng
		v := s[idx : idx+1]
		turn := k / leng
		cnt := bits.OnesCount(uint(turn))
		if cnt%2 > 0 {
			v = sre[idx : idx+1]
		}

		ans[i] = v
	}

	str := fmt.Sprintf("%v", ans)
	fmt.Fprint(out, str[1:len(str)-1])
}

func reverseCase(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			result[i] = unicode.ToUpper(r)
		}
	}
	return string(result)
}

func main() {
	run(os.Stdin, os.Stdout)
}
