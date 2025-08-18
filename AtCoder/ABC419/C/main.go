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
	var nodes [][]int
	for i := 0; i < n; i++ {
		r, c := readInt(), readInt()
		if len(nodes) == 0 {
			nodes = make([][]int, 4)
			for idx := range nodes {
				nodes[idx] = []int{r, c}
			}
			continue
		}

		newnode := []int{r, c}
		if nodes[0][0] > r {
			nodes[0] = newnode
		}
		if nodes[1][0] < r {
			nodes[1] = newnode
		}
		if nodes[2][1] > c {
			nodes[2] = newnode
		}
		if nodes[3][1] < c {
			nodes[3] = newnode
		}
	}

	ml := 0
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			dis := max(abs(nodes[i][0]-nodes[j][0]), abs(nodes[i][1]-nodes[j][1]))
			ml = max(ml, dis)
		}
	}

	fmt.Fprint(out, (ml+1)/2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	run(os.Stdin, os.Stdout)
}
