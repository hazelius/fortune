package main

import "testing"

func Test_run(t *testing.T) {
	mm := 200000
	mxy := make([][]int, mm)
	for i := range mxy {
		mxy[i] = []int{i, i}
	}

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
		{name: "1", args: args{n: 2, m: 4, xys: [][]int{
			{1, 1},
			{1, 2},
			{2, 0},
			{4, 2},
		}}, want: 3},
		{name: "2", args: args{n: 1, m: 1, xys: [][]int{
			{1, 1},
		}}, want: 0},
		{name: "3", args: args{n: 4, m: 72, xys: [][]int{
			{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0},
			{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}, {6, 1}, {7, 1}, {8, 1},
			{1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}, {8, 2},
			{1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {6, 3}, {7, 3}, {8, 3},
			{1, 4}, {2, 4}, {3, 4}, {4, 4}, {5, 4}, {6, 4}, {7, 4}, {8, 4},
			{1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {6, 5}, {7, 5}, {8, 5},
			{1, 6}, {2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}, {7, 6}, {8, 6},
			{1, 7}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7}, {8, 7},
			{1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8}, {6, 8}, {7, 8}, {8, 8},
		}}, want: 5},
		{name: "4", args: args{n: 100000000, m: mm, xys: mxy}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.m, tt.args.xys); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
