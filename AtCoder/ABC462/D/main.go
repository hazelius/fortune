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
	buf := make([]byte, 1<<20)
	sc.Buffer(buf, len(buf))
	sc.Split(bufio.ScanWords)

	n, d := readInt(), readInt()
	smap := make(map[int]int)
	tmap := make(map[int]int)
	maxt := 0
	for i := 0; i < n; i++ {
		s, t := readInt(), readInt()-d+1
		if s >= t {
			continue
		}
		smap[s]++
		tmap[t]++
		if maxt < t {
			maxt = t
		}
	}

	ans := 0
	cnt := 0
	for i := 1; i < maxt; i++ {
		if s, ok := smap[i]; ok {
			cnt += s
		}
		if t, ok := tmap[i]; ok {
			cnt -= t
		}
		if cnt >= 2 {
			ans += cnt * (cnt - 1) / 2
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
