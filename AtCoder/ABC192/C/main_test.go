package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 314, k: 2}, want: "693"},
		{name: "2", args: args{n: 1000000000, k: 100}, want: "0"},
		{name: "3", args: args{n: 6174, k: 100000}, want: "6174"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
