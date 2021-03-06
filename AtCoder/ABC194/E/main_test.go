package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		r []int
		m int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: []int{0, 0, 1}, m: 2}, want: "1"},
		{name: "2", args: args{r: []int{1, 1, 1}, m: 2}, want: "0"},
		{name: "3", args: args{r: []int{0, 1, 0}, m: 2}, want: "2"},
		{name: "4", args: args{r: []int{0, 0, 1, 2, 0, 1, 0}, m: 3}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r, tt.args.m); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
