package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		y string
		m uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{y: "22", m: 10}, want: "2"},
		{name: "2", args: args{y: "999", m: 1500}, want: "3"},
		{name: "3", args: args{y: "100000000000000000000000000000000000000000000000000000000000", m: 1000000000000000000}, want: "1"},
		{name: "4", args: args{y: "11", m: 99000000}, want: "98999998"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.y, tt.args.m); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
