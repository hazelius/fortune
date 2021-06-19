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
		{name: "1", args: args{n: 8, as: []int{1, 5, 3, 2, 5, 2, 3, 1}}, want: 2},
		{name: "2", args: args{n: 7, as: []int{1, 2, 3, 4, 1, 2, 3}}, want: 1},
		{name: "3", args: args{n: 1, as: []int{200000}}, want: 0},
		{name: "4", args: args{n: 8, as: []int{
			2, 3, 1, 1,
			1, 3, 1, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
