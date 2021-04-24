package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n    int
		s    string
		tabs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 2, s: "FLIP", tabs: [][]int{{2, 0, 0}, {1, 1, 4}}}, want: "LPFI"},
		{name: "2", args: args{n: 2, s: "FLIP", tabs: [][]int{
			{1, 1, 3},
			{2, 0, 0},
			{1, 1, 2},
			{1, 2, 3},
			{2, 0, 0},
			{1, 1, 4},
		}}, want: "ILPF"},
		{name: "3", args: args{n: 5, s: "ABCDEFGHIJ", tabs: [][]int{
			{2, 0, 0},
			{1, 1, 10},
			{1, 5, 6},
			{1, 7, 8},
			{2, 0, 0},
			{1, 8, 8},
			{1, 6, 8},
			{2, 0, 0},
		}}, want: "HGEIAJCBDF"},
		{name: "4", args: args{n: 1, s: "AB", tabs: [][]int{
			{1, 1, 2},
			{2, 0, 0},
			{2, 0, 0},
		}}, want: "BA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.s, tt.args.tabs); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
