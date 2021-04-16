// L - Deque
// https://atcoder.jp/contests/dp/tasks/dp_l

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
		{name: "1", args: args{n: 4, as: []int{10, 80, 90, 30}}, want: 10},
		{name: "2", args: args{n: 3, as: []int{10, 100, 10}}, want: -80},
		{name: "3", args: args{n: 1, as: []int{10}}, want: 10},
		{name: "4", args: args{n: 10, as: []int{1000000000, 1, 1000000000, 1, 1000000000, 1, 1000000000, 1, 1000000000, 1}}, want: 4999999995},
		{name: "5", args: args{n: 6, as: []int{4, 2, 9, 7, 1, 5}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
