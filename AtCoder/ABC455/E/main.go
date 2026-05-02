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

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n := readInt()
	s := readString()
	ans := n * (n + 1) / 2

	aminusbmap := make(map[int]int)
	aminuscmap := make(map[int]int)
	bminuscmap := make(map[int]int)
	abcmap := make(map[[2]int]int)
	aminusbmap[0] = 1
	aminuscmap[0] = 1
	bminuscmap[0] = 1
	abcmap[[2]int{0, 0}] = 1

	aminusb := 0
	aminusc := 0
	bminusc := 0
	for _, c := range s {
		switch c {
		case 'A':
			aminusb++
			aminusc++
		case 'B':
			aminusb--
			bminusc++
		case 'C':
			aminusc--
			bminusc--
		}
		ans -= aminusbmap[aminusb]
		ans -= aminuscmap[aminusc]
		ans -= bminuscmap[bminusc]
		ans += 2 * abcmap[[2]int{aminusb, aminusc}]
		aminusbmap[aminusb]++
		aminuscmap[aminusc]++
		bminuscmap[bminusc]++
		abcmap[[2]int{aminusb, aminusc}]++
	}

	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
