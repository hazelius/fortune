package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1", args: args{n: 12}, want: 5},
		{name: "2", args: args{n: 100128}, want: 447},
		{name: "3", args: args{n: 0}, want: 1},
		{name: "4", args: args{n: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
