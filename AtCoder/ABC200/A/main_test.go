// https://atcoder.jp/contests/abc200/tasks/abc200_a
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 2021}, want: 21},
		{name: "2", args: args{n: 200}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
