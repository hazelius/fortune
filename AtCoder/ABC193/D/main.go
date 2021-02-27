package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func point(a, b int) int {
	r := a
	for i := 0; i < b; i++ {
		r *= 10
	}
	return r
}

func sum(m map[int]int) int {
	var r int
	for i := 1; i < 10; i++ {
		r += point(i, m[i])
	}
	return r
}

func winNum(s, t map[int]int) []int {
	result := make([]int, 0, 9)

	sSum := sum(s)
	for i := 1; i < 10; i++ {
		t[i]++
		if sSum > sum(t) {
			result = append(result, i)
		}
		t[i]--
	}

	return result
}

func winRate(nums []int, m map[int]int, nokori float64) float64 {
	wincard := 0
	for _, num := range nums {
		wincard += m[num]
	}
	// fmt.Printf("m:%v, wincard:%v, nokori:%v\n", m, wincard, nokori)
	return float64(wincard) / nokori
}

func run(k, s, t string) float64 {
	ks, _ := strconv.Atoi(k)
	m := make(map[int]int, 9)
	for i := 1; i <= 9; i++ {
		m[i] = ks
	}

	sm := make(map[int]int, 9)
	tm := make(map[int]int, 9)
	for i := 0; i < 4; i++ {
		si, _ := strconv.Atoi(s[i : i+1])
		sm[si]++
		ti, _ := strconv.Atoi(t[i : i+1])
		tm[ti]++

		m[si]--
		m[ti]--
	}

	nokoriS := float64(ks*9 - 8)
	nokoriT := float64(ks*9 - 9)

	var result float64
	for i := 1; i < 10; i++ {
		sm[i]++
		m[i]--

		nums := winNum(sm, tm)
		rate := winRate(nums, m, nokoriT)

		sm[i]--
		m[i]++

		result += rate * (float64(m[i]) / (nokoriS))
		// fmt.Printf("winnum:%v, i:%v rate:%v, result:%v\n", nums, i, rate, result)
	}

	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	fmt.Println(run(read(), read(), read()))
}
