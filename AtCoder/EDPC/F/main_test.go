// https://atcoder.jp/contests/dp/tasks/dp_f

package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "axyb", t: "abyxb"}, want: "axb"},
		{name: "2", args: args{s: "aa", t: "xayaz"}, want: "aa"},
		{name: "3", args: args{s: "a", t: "x"}, want: ""},
		{name: "4", args: args{s: "abracadabra", t: "avadakedavra"}, want: "aaadara"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
