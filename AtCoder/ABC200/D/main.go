package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(n int, as []int) string {
	dp := make([][]int, 200)
	if n > 8 {
		n = 8
	}
	for i := 1; i < 1<<n; i++ {
		x := 0
		arr := make([]int, 0)
		for j := 0; j < n; j++ {
			if i>>j&1 > 0 {
				x = (x + as[j]) % 200
				arr = append(arr, j+1)
			}
		}
		if len(dp[x]) > 0 {
			b := fmt.Sprintf("%v", dp[x])
			c := fmt.Sprintf("%v", arr)
			return fmt.Sprintf("Yes\n%v %v\n%v %v", len(dp[x]), b[1:len(b)-1], len(arr), c[1:len(c)-1])
		}
		dp[x] = arr
	}
	return "No"
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	as := make([]int, n)
	for i := range as {
		as[i] = readInt()
	}
	fmt.Println(run(n, as))
}
