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

	n := readInt()
	ks := make([]int, n)
	kas := make([]map[int]int, n)
	for i := range ks {
		ks[i] = readInt()
		kas[i] = make(map[int]int)
		for j := 0; j < ks[i]; j++ {
			a := readInt()
			kas[i][a] += 1
		}
	}

	ans := 0.0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			mapa := kas[i]
			mapb := kas[j]
			if len(mapa) > len(mapb) {
				mapa, mapb = mapb, mapa
			}

			sum := 0
			for a, cnt := range mapa {
				cnt2, ok := mapb[a]
				if ok {
					sum += cnt * cnt2
				}
			}
			d := float64(sum) / float64(ks[i]*ks[j])
			if d > ans {
				ans = d
			}
		}
	}

	fmt.Fprintf(out, "%.15f", ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
