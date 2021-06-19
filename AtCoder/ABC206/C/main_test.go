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
		{name: "1", args: args{n: 3, as: []int{1, 7, 1}}, want: 2},
		{name: "2", args: args{n: 10, as: []int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}}, want: 45},
		{name: "3", args: args{n: 20, as: []int{7, 8, 1, 1, 4, 9, 9, 6, 8, 2, 4, 1, 1, 9, 5, 5, 5, 3, 6, 4}}, want: 173},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
