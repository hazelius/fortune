package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "dIfFiCuLt"}, want: "Yes"},
		{name: "2", args: args{s: "eASY"}, want: "No"},
		{name: "3", args: args{s: "a"}, want: "Yes"},
		{name: "4", args: args{s: "EASY"}, want: "No"},
		{name: "5", args: args{s: "A"}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
