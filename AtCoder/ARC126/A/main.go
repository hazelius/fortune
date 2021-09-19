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

	t := readInt()

	answers := make([]int, t)
	for i := 0; i < t; i++ {
		n2, n3, n4 := readInt(), readInt(), readInt()
		ans := 0

		n3 = n3 / 2
		if n3 > n4 {
			ans += n4
			n3 -= n4

			if n3 < n2/2 {
				ans += n3
				n2 -= n3 * 2
				ans += n2 / 5
			} else {
				ans += n2 / 2
			}
		} else {
			ans += n3
			n4 -= n3

			n4amari := n4 % 2
			n4 = n4 / 2
			if n4 < n2 {
				ans += n4
				n2 -= n4

				if n4amari > 0 && n2 >= 3 {
					ans++
					n2 -= 3
				}
				ans += n2 / 5
			} else {
				ans += n2
			}
		}
		answers[i] = ans
	}

	ret := fmt.Sprintf("%v", answers)
	return ret[1 : len(ret)-1]
}

func main() {
	fmt.Println(run(os.Stdin))
}
