package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{a: 100, b: 80}, want: "20"},
		{name: "1", args: args{a: 7, b: 6}, want: "14.285714285714285714"},
		{name: "1", args: args{a: 99999, b: 99998}, want: "0.00100001000010000100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
