package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		abc [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, abc: [][]int{{10, 40, 70}, {20, 50, 80}, {30, 60, 90}}}, want: 210},
		{name: "2", args: args{n: 1, abc: [][]int{{100, 10, 1}}}, want: 100},
		{name: "3", args: args{n: 7, abc: [][]int{
			{6, 7, 8},
			{8, 8, 3},
			{2, 5, 2},
			{7, 8, 6},
			{4, 6, 8},
			{2, 3, 4},
			{7, 5, 1},
		}}, want: 46},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.abc); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
