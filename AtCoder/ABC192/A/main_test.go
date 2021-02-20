package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{x: 140}, want: "60"},
		{name: "2", args: args{x: 1000}, want: "100"},
		{name: "3", args: args{x: 1099}, want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.x); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
