package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		as []int
		bs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 2, as: []int{3, 2}, bs: []int{7, 5}}, want: 3},
		{name: "2", args: args{n: 3, as: []int{1, 5, 3}, bs: []int{10, 7, 3}}, want: 0},
		{name: "3", args: args{n: 3, as: []int{3, 2, 5}, bs: []int{6, 9, 8}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as, tt.args.bs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
