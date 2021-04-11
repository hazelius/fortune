package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

var (
	v1   string
	v2   string
	v3   string
	ans1 int
	ans2 int
	ans3 int
)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func run(s1, s2, s3 string) string {
	v1, v2, v3 = s1, s2, s3
	nokori := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	nums := make([]int, 10)

	ans := makeArray(nums, 0, nokori)
	if ans {
		return fmt.Sprintf("%v\n%v\n%v", ans1, ans2, ans3)
	}
	return "UNSOLVABLE"
}

func makeArray(nums []int, idx int, nokori []int) bool {
	for i, v := range nokori {
		nums[idx] = v
		if len(nokori) == 1 {
			if calc(nums) {
				return true
			}
		} else {
			tmp := make([]int, len(nokori)-1)
			tmpi := 0
			for copyi, v := range nokori {
				if copyi != i {
					tmp[tmpi] = v
					tmpi++
				}
			}
			if makeArray(nums, idx+1, tmp) {
				return true
			}
		}
	}
	return false
}

func calc(nums []int) bool {
	chars := make(map[rune]int)
	ci := 0
	v1int := 0
	for i, v := range v1 {
		n, ok := chars[v]
		if !ok {
			if ci >= len(nums) {
				return false
			}

			n = nums[ci]
			if i == 0 && n == 0 {
				return false
			}
			chars[v] = n
			ci++
		}
		v1int = 10*v1int + n
	}

	v2int := 0
	for i, v := range v2 {
		n, ok := chars[v]
		if !ok {
			if ci >= len(nums) {
				return false
			}

			n = nums[ci]
			if i == 0 && n == 0 {
				return false
			}
			chars[v] = n
			ci++
		}
		v2int = 10*v2int + n
	}

	v3int := 0
	for i, v := range v3 {
		n, ok := chars[v]
		if !ok {
			if ci >= len(nums) {
				return false
			}

			n = nums[ci]
			if i == 0 && n == 0 {
				return false
			}
			chars[v] = n
			ci++
		}
		v3int = 10*v3int + n
	}

	if v1int+v2int == v3int {
		ans1 = v1int
		ans2 = v2int
		ans3 = v3int
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	s1, s2, s3 := readString(), readString(), readString()
	fmt.Println(run(s1, s2, s3))
}
