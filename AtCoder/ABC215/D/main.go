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

	n, m := readInt(), readInt()
	as := make([]bool, int(1e5+1))
	for i := 0; i < n; i++ {
		as[readInt()] = true
	}

	nums := make([]bool, int(1e5+1))
	nums[0], nums[1] = true, true
	pns := make([]int, 0)
	for i, b := range nums {
		if b {
			continue
		}
		pns = append(pns, i)
		for j := i * 2; j < len(nums); j += i {
			nums[j] = true
		}
	}

	used := make([]bool, m+1)
	for _, pn := range pns {
		if pn > m {
			break
		}
		flg := false
		for j := pn; j <= int(1e5); j += pn {
			if as[j] {
				flg = true
				break
			}
		}
		if !flg {
			continue
		}

		for j := pn; j <= m; j += pn {
			used[j] = true
		}
	}

	ans := make([]int, 0)
	for i := 1; i <= m; i++ {
		if !used[i] {
			ans = append(ans, i)
		}
	}

	ansstring := fmt.Sprintf("%v", ans)
	return fmt.Sprintf("%v %v", len(ans), ansstring[1:len(ansstring)-1])
}

func main() {
	fmt.Println(run(os.Stdin))
}
