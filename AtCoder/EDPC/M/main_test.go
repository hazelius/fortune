// M - Candies
// https://atcoder.jp/contests/dp/tasks/dp_m

package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		k  int
		as []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, k: 4, as: []int{1, 2, 3}}, want: 5},
		{name: "2", args: args{n: 1, k: 10, as: []int{9}}, want: 0},
		{name: "3", args: args{n: 2, k: 0, as: []int{0, 0}}, want: 1},
		// {name: "4", args: args{n: 4, k: 100000, as: []int{100000, 100000, 100000, 100000}}, want: 665683269},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
