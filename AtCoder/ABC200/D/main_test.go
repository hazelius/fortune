package main

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		n  int
		as []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{n: 5, as: []int{180, 186, 189, 191, 218}}, want: "Yes\n1 1\n2 3 4"},
		{name: "2", args: args{n: 2, as: []int{123, 523}}, want: "Yes\n1 1\n1 2"},
		{name: "3", args: args{n: 6, as: []int{2013, 1012, 2765, 2021, 508, 6971}}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n, tt.args.as); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
