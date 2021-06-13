package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n  int
		q  int
		as []int
		ks []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{n: 4, q: 3, as: []int{3, 5, 6, 7}, ks: []int{2, 5, 3}}, want: []int{2, 9, 4}},
		{name: "2", args: args{n: 5, q: 2, as: []int{1, 2, 3, 4, 5}, ks: []int{1, 10}}, want: []int{6, 15}},
		{name: "3", args: args{n: 5, q: 2, as: []int{1, 10, 15, 34, 28}, ks: []int{1, 10}}, want: []int{2, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.q, tt.args.as, tt.args.ks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
