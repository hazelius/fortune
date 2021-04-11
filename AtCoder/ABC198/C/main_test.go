package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		r int
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{r: 5, x: 15, y: 0}, want: 3},
		{name: "2", args: args{r: 5, x: 11, y: 0}, want: 3},
		{name: "3", args: args{r: 3, x: 4, y: 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
