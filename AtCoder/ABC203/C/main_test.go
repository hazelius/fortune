package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		k   int
		abs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 2, k: 3, abs: [][]int{{2, 1}, {5, 10}}}, want: 4},
		{name: "2", args: args{n: 5, k: 1000000000, abs: [][]int{
			{1, 1000000000},
			{2, 1000000000},
			{3, 1000000000},
			{4, 1000000000},
			{5, 1000000000},
		}}, want: 6000000000},
		{name: "3", args: args{n: 3, k: 2, abs: [][]int{
			{5, 5},
			{2, 1},
			{2, 2},
		}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.abs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
