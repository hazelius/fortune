package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{a: 45, b: 200}, want: 90},
		{name: "2", args: args{a: 37, b: 450}, want: 166.5},
		{name: "3", args: args{a: 0, b: 1000}, want: 0},
		{name: "4", args: args{a: 50, b: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
