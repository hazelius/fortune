package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	_, q := readInt(), readInt()
	front := make(map[int]int)
	back := make(map[int]int)
	ans := make([]string, 0)
	for i := 0; i < q; i++ {
		switch readInt() {
		case 1:
			x, y := readInt(), readInt()
			front[y] = x
			back[x] = y
		case 2:
			x, y := readInt(), readInt()
			delete(front, y)
			delete(back, x)
		case 3:
			x := readInt()
			for {
				v, ok := front[x]
				if !ok {
					break
				}
				x = v
			}

			train := []int{x}
			for {
				v, ok := back[x]
				if !ok {
					break
				}
				train = append(train, v)
				x = v
			}

			train = append([]int{len(train)}, train...)
			trainStr := fmt.Sprintf("%v", train)
			ans = append(ans, trainStr[1:len(trainStr)-1])
		}
	}
	return strings.Join(ans, "\n")
}

func main() {
	fmt.Println(run(os.Stdin))
}
