package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "erasedream"}, want: "YES"},
		{name: "2", args: args{s: "dreameraser"}, want: "YES"},
		{name: "3", args: args{s: "dreamerer"}, want: "NO"},
		{name: "4", args: args{s: "dream"}, want: "YES"},
		{name: "4", args: args{s: "dreamerererase"}, want: "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
