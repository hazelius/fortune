// K - Stones
// https://atcoder.jp/contests/dp/tasks/dp_k

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
		want string
	}{
		{name: "1", args: args{n: 2, k: 4, as: []int{2, 3}}, want: "First"},
		{name: "2", args: args{n: 2, k: 5, as: []int{2, 3}}, want: "Second"},
		{name: "3", args: args{n: 2, k: 7, as: []int{2, 3}}, want: "First"},
		{name: "4", args: args{n: 3, k: 20, as: []int{1, 2, 3}}, want: "Second"},
		{name: "5", args: args{n: 3, k: 21, as: []int{1, 2, 3}}, want: "First"},
		{name: "6", args: args{n: 1, k: 100000, as: []int{1}}, want: "Second"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
