package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int
		k int
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 5, k: 3, h: []int{10, 30, 40, 50, 20}}, want: 30},
		{name: "2", args: args{n: 3, k: 1, h: []int{10, 20, 10}}, want: 20},
		{name: "3", args: args{n: 2, k: 100, h: []int{10, 10}}, want: 0},
		{name: "4", args: args{n: 10, k: 4, h: []int{40, 10, 20, 70, 80, 10, 20, 70, 80, 60}}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.h); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
