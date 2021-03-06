package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		r [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: [][]int{{8, 5}, {4, 4}, {7, 9}}}, want: "5"},
		{name: "2", args: args{r: [][]int{{11, 7}, {3, 2}, {6, 7}}}, want: "5"},
		{name: "3", args: args{r: [][]int{{11, 7}, {3, 3}, {3, 7}}}, want: "3"},
		{name: "4", args: args{r: [][]int{{11, 7}}}, want: "18"},
		{name: "5", args: args{r: [][]int{{11, 7}, {11, 7}}}, want: "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
