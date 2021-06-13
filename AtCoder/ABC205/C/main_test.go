package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a: 3, b: 2, c: 4}, want: ">"},
		{name: "2", args: args{a: -7, b: 7, c: 2}, want: "="},
		{name: "3", args: args{a: -8, b: 6, c: 3}, want: "<"},
		{name: "4", args: args{a: -8, b: 9, c: 3}, want: "<"},
		{name: "5", args: args{a: -8, b: 9, c: 4}, want: "<"},
		{name: "6", args: args{a: -9, b: 9, c: 4}, want: "="},
		{name: "7", args: args{a: -10, b: 9, c: 4}, want: ">"},
		{name: "7", args: args{a: -10, b: 9, c: 5}, want: "<"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
