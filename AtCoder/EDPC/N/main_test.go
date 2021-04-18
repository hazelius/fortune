// N - Slimes
// https://atcoder.jp/contests/dp/tasks/dp_n
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		as []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{n: 3, as: []int{10, 20, 20}}, want: 80},
		{name: "1", args: args{n: 4, as: []int{10, 20, 30, 40}}, want: 190},
		{name: "2", args: args{n: 5, as: []int{10, 10, 10, 10, 10}}, want: 120},
		{name: "3", args: args{n: 3, as: []int{1000000000, 1000000000, 1000000000}}, want: 5000000000},
		{name: "4", args: args{n: 6, as: []int{7, 6, 8, 6, 1, 1}}, want: 68},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
