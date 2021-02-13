package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		x int
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{x: 5, a: []int{3, 5, 6, 5, 5, 4}}, want: "3 6 4"},
		{name: "2", args: args{x: 3, a: []int{3, 3, 3}}, want: ""},
		{name: "3", args: args{x: 2, a: []int{3, 1, 3, 2, 2, 3, 2, 3, 4}}, want: "3 1 3 3 3 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.x, tt.args.a); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
