package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		m int
		h int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{m: 10, h: 120}, want: "Yes"},
		{name: "2", args: args{m: 10, h: 125}, want: "No"},
		{name: "2", args: args{m: 1, h: 1}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.m, tt.args.h); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
