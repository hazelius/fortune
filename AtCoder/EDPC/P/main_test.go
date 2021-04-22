// P - Independent Set
// https://atcoder.jp/contests/dp/tasks/dp_p

package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		xys [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, xys: [][]int{{1, 2}, {2, 3}}}, want: 5},
		{name: "2", args: args{n: 4, xys: [][]int{{1, 2}, {1, 3}, {1, 4}}}, want: 9},
		{name: "3", args: args{n: 1, xys: [][]int{}}, want: 2},
		{name: "4", args: args{n: 10, xys: [][]int{
			{8, 5},
			{10, 8},
			{6, 5},
			{1, 5},
			{4, 8},
			{2, 10},
			{3, 6},
			{9, 2},
			{1, 7},
		}}, want: 157},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.xys); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
