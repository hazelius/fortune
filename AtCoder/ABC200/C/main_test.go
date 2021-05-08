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
		{name: "1", args: args{n: 3, as: []int{123, 223, 123, 523, 200, 2000}}, want: 4},
		{name: "2", args: args{n: 5, as: []int{1, 2, 3, 4, 5}}, want: 0},
		{name: "3", args: args{n: 8, as: []int{199, 100, 200, 400, 300, 500, 600, 200}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
