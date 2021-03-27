package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  float64
		x0 float64
		y0 float64
		xn float64
		yn float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 4, x0: 1, y0: 1, xn: 2, yn: 2}, want: "2.0000000 1.0000000"},
		{name: "2", args: args{n: 6, x0: 5, y0: 3, xn: 7, yn: 4}, want: "5.9330127 2.3839746"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.x0, tt.args.y0, tt.args.xn, tt.args.yn); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
