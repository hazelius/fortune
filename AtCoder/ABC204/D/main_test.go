package main

import "testing"

func Test_run(t *testing.T) {
	n := 100
	ts := make([]int, n)
	for i := range ts {
		ts[i] = i + 1
	}

	type args struct {
		n  int
		ts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 5, ts: []int{8, 3, 7, 2, 5}}, want: 13},
		{name: "2", args: args{n: 2, ts: []int{1000, 1}}, want: 1000},
		{name: "3", args: args{n: 9, ts: []int{3, 14, 15, 9, 26, 5, 35, 89, 79}}, want: 138},
		{name: "4", args: args{n: 6, ts: []int{100, 100, 14, 15, 210, 420}}, want: 434},
		{name: "5", args: args{n: 2, ts: []int{1, 1000}}, want: 1000},
		{name: "6", args: args{n: 3, ts: []int{1, 5, 1}}, want: 5},
		{name: "7", args: args{n: n, ts: ts}, want: 2525},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.ts); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
