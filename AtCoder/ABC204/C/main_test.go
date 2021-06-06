package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		m  int
		ab [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, m: 3, ab: [][]int{
			{1, 2},
			{2, 3},
			{3, 2},
		}}, want: 7},
		{name: "2", args: args{n: 3, m: 0, ab: [][]int{}}, want: 3},
		{name: "3", args: args{n: 4, m: 4, ab: [][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{4, 1},
		}}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.m, tt.args.ab); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
