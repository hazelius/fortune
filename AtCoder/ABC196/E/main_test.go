package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		fs [][]int
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{fs: [][]int{{-10, 2}, {10, 1}, {10, 3}}, xs: []int{-15, -10, -5, 0, 5}}, want: []int{0, 0, 5, 10, 10}},
		{name: "2", args: args{
			fs: [][]int{{-10, 1}, {-10, 1}, {-20, 2}, {10, 2}, {40, 3}, {30, 3}},
			xs: []int{0, 30, -100}},
			want: []int{10, 10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.fs, tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
