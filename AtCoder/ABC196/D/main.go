package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var gh, gw, ans int

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(h, w, a, b int) int {
	gh, gw, ans = h, w, 0
	bit := make([]bool, h*w)
	dfs(0, bit, a, b)
	return ans
}

func dfs(i int, bit []bool, a, b int) {
	if i == gh*gw-1 {
		ans++
		// fmt.Printf("%v,%v,%v\n", i, bit, ans)
		return
	}
	if bit[i] {
		// fmt.Printf("%v,%v,%v\n", i, bit, ans)
		dfs(i+1, bit, a, b)
		return
	}
	if b > 0 {
		newbit := make([]bool, len(bit))
		copy(newbit, bit)
		newbit[i] = true
		dfs(i+1, newbit, a, b-1)
	}
	if a > 0 {
		if i%gw != gw-1 && !bit[i+1] {
			newbit := make([]bool, len(bit))
			copy(newbit, bit)
			newbit[i] = true
			newbit[i+1] = true
			dfs(i+1, newbit, a-1, b)
		}
		if i+gw < gh*gw && !bit[i+gw] {
			newbit := make([]bool, len(bit))
			copy(newbit, bit)
			newbit[i] = true
			newbit[i+gw] = true
			dfs(i+1, newbit, a-1, b)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w, a, b := readInt(), readInt(), readInt(), readInt()
	fmt.Println(run(h, w, a, b))
}
