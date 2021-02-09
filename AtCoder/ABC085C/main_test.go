package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 9, y: 45000}, want: "0 9 0"},
		{name: "2", args: args{n: 1000, y: 1234000}, want: "2 54 944"},
		{name: "3", args: args{n: 2000, y: 20000000}, want: "2000 0 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.y); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
