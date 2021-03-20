package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		h int
		w int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{h: 2, w: 2, a: 1, b: 2}, want: 4},
		{name: "2", args: args{h: 3, w: 3, a: 4, b: 1}, want: 18},
		{name: "3", args: args{h: 4, w: 4, a: 8, b: 0}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.h, tt.args.w, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
