package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 180}, want: "Yay!"},
		{name: "2", args: args{n: 200}, want: ":("},
		{name: "3", args: args{n: 191}, want: "so-so"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
