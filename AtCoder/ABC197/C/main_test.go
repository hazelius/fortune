package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{arr: []int{1, 5, 7}}, want: 2},
		{name: "2", args: args{arr: []int{10, 10, 10}}, want: 0},
		{name: "3", args: args{arr: []int{1, 3, 3, 1}}, want: 0},
		{name: "4", args: args{arr: []int{1, 3, 3, 1, 4, 4, 7, 10, 10}}, want: 0},
		{name: "5", args: args{arr: []int{15, 1, 16, 15, 17}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
