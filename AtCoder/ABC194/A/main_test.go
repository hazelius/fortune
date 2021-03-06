package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a: 10, b: 8}, want: "1"},
		{name: "2", args: args{a: 1, b: 2}, want: "3"},
		{name: "3", args: args{a: 0, b: 0}, want: "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
