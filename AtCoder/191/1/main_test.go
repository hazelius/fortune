package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		v int
		t int
		s int
		d int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{v: 10, t: 3, s: 5, d: 20}, want: "Yes"},
		{name: "2", args: args{v: 10, t: 3, s: 5, d: 30}, want: "No"},
		{name: "3", args: args{v: 15, t: 5, s: 7, d: 105}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.v, tt.args.t, tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
