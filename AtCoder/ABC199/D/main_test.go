package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		m   int
		abs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, m: 3, abs: [][]int{{1, 2}, {2, 3}, {3, 1}}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.m, tt.args.abs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
