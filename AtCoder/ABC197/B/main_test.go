package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		h  int
		w  int
		x  int
		y  int
		ss []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{h: 4, w: 4, x: 2, y: 2, ss: []string{
			"##..",
			"...#",
			"#.#.",
			".#.#",
		}}, want: 4},
		{name: "2", args: args{h: 3, w: 5, x: 1, y: 4, ss: []string{
			"#....",
			"#####",
			"....#",
		}}, want: 4},
		{name: "3", args: args{h: 5, w: 5, x: 4, y: 2, ss: []string{
			".#..#",
			"#.###",
			"##...",
			"#..#.",
			"#.###",
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.h, tt.args.w, tt.args.x, tt.args.y, tt.args.ss); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
