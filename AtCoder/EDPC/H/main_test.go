// https://atcoder.jp/contests/dp/tasks/dp_h
package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		h   int
		w   int
		ahw []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{h: 3, w: 4, ahw: []string{
			"...#",
			".#..",
			"....",
		}}, want: 3},
		{name: "2", args: args{h: 5, w: 2, ahw: []string{
			"..",
			"#.",
			"..",
			".#",
			"..",
		}}, want: 0},
		{name: "3", args: args{h: 5, w: 5, ahw: []string{
			"..#..",
			".....",
			"#...#",
			".....",
			"..#..",
		}}, want: 24},
		{name: "4", args: args{h: 20, w: 20, ahw: []string{
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
			"....................",
		}}, want: 345263555},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.h, tt.args.w, tt.args.ahw); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
