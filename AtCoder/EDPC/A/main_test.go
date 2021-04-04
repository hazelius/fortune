package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{4, []int{10, 30, 40, 20}}, want: 30},
		{name: "2", args: args{2, []int{10, 10}}, want: 0},
		{name: "3", args: args{6, []int{30, 10, 60, 10, 60, 50}}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.h); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
