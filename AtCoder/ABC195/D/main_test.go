package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n [][]int
		m []int
		q [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{
			n: [][]int{{1, 9}, {5, 3}, {7, 8}},
			m: []int{1, 8, 6, 9},
			q: [][]int{{4, 4}, {1, 4}, {1, 3}},
		}, want: []int{20, 0, 9}},
		{name: "1", args: args{
			n: [][]int{{1, 9}, {5, 3}, {7, 8}, {6, 3}, {8, 9}, {2, 1}},
			m: []int{1, 8, 6, 9, 3, 3},
			q: [][]int{{5, 5}, {1, 4}, {2, 4}},
		}, want: []int{30, 10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.m, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
