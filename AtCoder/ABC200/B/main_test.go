package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 2021, k: 4}, want: 50531},
		{name: "2", args: args{n: 40000, k: 2}, want: 1},
		{name: "3", args: args{n: 8691, k: 20}, want: 84875488281},
		{name: "4", args: args{n: 2021, k: 5}, want: 50531200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
