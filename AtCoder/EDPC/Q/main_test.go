// Q - Flowers
// https://atcoder.jp/contests/dp/tasks/dp_q
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		hs []int
		as []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 4, hs: []int{3, 1, 4, 2}, as: []int{10, 20, 30, 40}}, want: 60},
		{name: "2", args: args{n: 1, hs: []int{1}, as: []int{10}}, want: 10},
		{name: "3", args: args{n: 5, hs: []int{1, 2, 3, 4, 5}, as: []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}}, want: 5000000000},
		{name: "4", args: args{n: 9, hs: []int{4, 2, 5, 8, 3, 6, 1, 7, 9}, as: []int{6, 8, 8, 4, 6, 3, 5, 7, 5}}, want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.hs, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
