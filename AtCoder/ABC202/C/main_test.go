package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		as []int
		bs []int
		cs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 3, as: []int{1, 2, 2}, bs: []int{3, 1, 2}, cs: []int{2, 3, 2}}, want: 4},
		{name: "2", args: args{n: 4, as: []int{1, 1, 1, 1}, bs: []int{1, 1, 1, 1}, cs: []int{1, 2, 3, 4}}, want: 16},
		{name: "3", args: args{n: 3, as: []int{2, 3, 3}, bs: []int{1, 3, 3}, cs: []int{1, 1, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as, tt.args.bs, tt.args.cs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
