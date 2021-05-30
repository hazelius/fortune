package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		k  int
		sq [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, k: 2, sq: [][]int{
			{1, 7, 0},
			{5, 8, 11},
			{10, 4, 2},
		}}, want: 4},
		{name: "2", args: args{n: 3, k: 3, sq: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.sq); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
