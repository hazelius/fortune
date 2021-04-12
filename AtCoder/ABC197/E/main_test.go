package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		xcs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "1", args: args{xcs: [][]int{{2, 2}, {3, 1}, {1, 3}, {4, 2}, {5, 3}}}, want: 12},
		{name: "2", args: args{xcs: [][]int{
			{5, 5},
			{-4, 4},
			{4, 3},
			{6, 3},
			{-5, 5},
			{-3, 2},
			{2, 2},
			{3, 3},
			{1, 4},
		}}, want: 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.xcs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
