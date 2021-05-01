package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n   int
		d   int
		h   int
		dhs [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{n: 1, d: 10, h: 10, dhs: [][]int{{3, 5}}}, want: 2.8571428571428568},
		{name: "2", args: args{n: 1, d: 10, h: 10, dhs: [][]int{{3, 2}}}, want: 0},
		{name: "3", args: args{n: 5, d: 896, h: 483, dhs: [][]int{
			{228, 59},
			{529, 310},
			{339, 60},
			{78, 266},
			{659, 391},
		}}, want: 245.30806845965773},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.d, tt.args.h, tt.args.dhs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
