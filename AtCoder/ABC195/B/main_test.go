package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a int
		b int
		w int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {name: "1", args: args{a: 100, b: 200, w: 2}, want: "10 20"},
		// {name: "2", args: args{a: 120, b: 150, w: 2}, want: "14 16"},
		{name: "3", args: args{a: 300, b: 333, w: 1}, want: "UNSATISFIABLE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.a, tt.args.b, tt.args.w); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
