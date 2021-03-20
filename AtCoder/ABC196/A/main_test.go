package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: 0, b: 10, c: 0, d: 10}, want: 10},
		{name: "2", args: args{a: -100, b: -100, c: 100, d: 100}, want: -200},
		{name: "3", args: args{a: -100, b: 100, c: -100, d: 100}, want: 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
