// https://atcoder.jp/contests/abc286/tasks/abc286_e
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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}

	pass := make([][][]int, n)
	for i := range pass {
		pass[i] = make([][]int, n)
		s := readString()
		for j, c := range s {
			if c == 'Y' {
				pass[i][j] = []int{1, as[j]}
			} else {
				pass[i][j] = []int{int(1e9 + 1), 0}
			}
		}
	}

	// ワーシャルフロイド法
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				root, price := pass[j][k][0], pass[j][k][1]
				rootA, priceA := pass[j][i][0], pass[j][i][1]
				rootB, priceB := pass[i][k][0], pass[i][k][1]
				if root > rootA+rootB {
					pass[j][k] = []int{rootA + rootB, priceA + priceB}
				} else if root == rootA+rootB && price < priceA+priceB {
					pass[j][k][1] = priceA + priceB
				}
			}
		}
	}

	q := readInt()
	for i := 0; i < q; i++ {
		u, v := readInt()-1, readInt()-1
		p := pass[u][v]
		if p[1] == 0 {
			fmt.Fprintln(out, "Impossible")
		} else {
			fmt.Fprintf(out, "%v %v\n", p[0], as[u]+p[1])
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
