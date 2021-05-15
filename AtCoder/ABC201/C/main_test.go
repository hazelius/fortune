package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "ooo???xxxx"}, want: 108},
		{name: "2", args: args{s: "o?oo?oxoxo"}, want: 0},
		{name: "3", args: args{s: "xxxxx?xxxo"}, want: 15},
		{name: "4", args: args{s: "xxxxx?xxx?"}, want: 16},
		{name: "5", args: args{s: "xxxxxoxxxo"}, want: 14},
		{name: "6", args: args{s: "xxxxxoooxo"}, want: 24},
		{name: "7", args: args{s: "ooo?xxxxxx"}, want: 60},
		{name: "8", args: args{s: "ooooxxxxx?"}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
