package main

import "testing"

func Test_run(t *testing.T) {
	as := make([]int, 1000)
	for i := range as {
		as[i] = 1000 - i
	}

	type args struct {
		n  int
		as []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 5, as: []int{3, 1, 2, 4, 5}}, want: "Yes"},
		{name: "2", args: args{n: 6, as: []int{3, 1, 4, 1, 5, 2}}, want: "No"},
		{name: "3", args: args{n: 3, as: []int{1, 2, 3}}, want: "Yes"},
		{name: "4", args: args{n: 1, as: []int{1}}, want: "Yes"},
		{name: "5", args: args{n: 1000, as: as}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
