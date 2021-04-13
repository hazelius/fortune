package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		ps []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{n: 3, ps: []float64{0.30, 0.60, 0.80}}, want: 0.612},
		{name: "2", args: args{n: 1, ps: []float64{0.50}}, want: 0.5},
		{name: "3", args: args{n: 5, ps: []float64{0.42, 0.01, 0.42, 0.99, 0.42}}, want: 0.3821815872},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.ps); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
