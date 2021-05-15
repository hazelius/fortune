// https://atcoder.jp/contests/abc200/tasks/abc200_a
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a1 int
		a2 int
		a3 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a1: 5, a2: 1, a3: 3}, want: "Yes"},
		{name: "2", args: args{a1: 1, a2: 4, a3: 3}, want: "No"},
		{name: "3", args: args{a1: 5, a2: 5, a3: 5}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a1, tt.args.a2, tt.args.a3); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
