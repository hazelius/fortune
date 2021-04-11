package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s1: "a", s2: "b", s3: "c"}, want: "1\n2\n3"},
		{name: "2", args: args{s1: "x", s2: "x", s3: "y"}, want: "1\n1\n2"},
		{name: "3", args: args{s1: "p", s2: "q", s3: "p"}, want: "UNSOLVABLE"},
		{name: "4", args: args{s1: "abcd", s2: "efgh", s3: "ijkl"}, want: "UNSOLVABLE"},
		{name: "5", args: args{s1: "send", s2: "more", s3: "money"}, want: "9567\n1085\n10652"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
