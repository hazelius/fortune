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
		{name: "1", args: args{n: 3, as: []int{6, 17, 28}}, want: 25},
		{name: "2", args: args{n: 4, as: []int{8, 9, 10, 11}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
