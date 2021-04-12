// https://atcoder.jp/contests/dp/tasks/dp_g
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		m   int
		xys [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 4, m: 5, xys: [][]int{{1, 2}, {1, 3}, {3, 2}, {2, 4}, {3, 4}}}, want: 3},
		{name: "2", args: args{n: 6, m: 3, xys: [][]int{{2, 3}, {4, 5}, {5, 6}}}, want: 2},
		{name: "3", args: args{n: 5, m: 8, xys: [][]int{
			{5, 3},
			{2, 3},
			{2, 4},
			{5, 2},
			{5, 1},
			{1, 4},
			{4, 3},
			{1, 3},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.m, tt.args.xys); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
