package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const MOD = 998244353

var fact []int64
var invFact []int64

func power(a, b int64) int64 {
	res := int64(1)
	a %= MOD
	for b > 0 {
		if b%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b /= 2
	}
	return res
}

func inverse(n int64) int64 {
	return power(n, MOD-2)
}

func precompute(n int) {
	fact = make([]int64, n+1)
	invFact = make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}
	invFact[n] = inverse(fact[n])
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}
}

func nCr(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	num := fact[n]
	den := invFact[r] * invFact[n-r] % MOD
	return num * den % MOD
}

var sc *bufio.Scanner

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(stdin io.Reader, out io.Writer) {
	sc = bufio.NewScanner(stdin)
	sc.Split(bufio.ScanWords)

	x1, x2, x3 := readInt(), readInt(), readInt()

	maxN := x2 + x3 + 5
	if x1 > maxN {
		maxN = x1 + 5
	}
	precompute(maxN)

	// この問題は、1と3が隣接してはいけないという制約のもとで、
	// 1のブロック(1-block)と3のブロック(3-block)を2によって区切る組み合わせを数える問題です。
	//
	// k個の非空の1-blockを作る方法は nCr(X1-1, k-1) 通りです。
	// m個の非空の3-blockを作る方法は nCr(X3-1, m-1) 通りです。
	// 合計 N = k+m 個のブロックを並める方法は nCr(k+m, k) 通りです。
	// 各ブロックの間には少なくとも1つの2が必要なため、N-1個の隙間に2を配置します。
	// 余った2を両端および隙間に配置する方法は、重複組合せの考え方で nCr(X2 - (N-1) + (N+1) - 1, (N+1) - 1) = nCr(X2+1, N) 通りです。
	//
	// したがって、全組み合わせは Σ_k Σ_m nCr(X1-1, k-1) * nCr(X3-1, m-1) * nCr(k+m, k) * nCr(X2+1, k+m) となります。
	// これを整理すると (ヴァンデルモンドの等式を利用):
	// Σ_k nCr(X1-1, k-1) * nCr(X2+1, k) * nCr(X2+X3-k, X3) となり、O(X1) で計算可能です。

	ans := int64(0)
	for k := 1; k <= x1; k++ {
		term := nCr(x2+1, k)
		if term == 0 {
			continue
		}
		term = term * nCr(x1-1, k-1) % MOD
		term = term * nCr(x2+x3-k, x3) % MOD
		ans = (ans + term) % MOD
	}

	fmt.Fprintln(out, ans)
}

func main() {
	run(os.Stdin, os.Stdout)
}
