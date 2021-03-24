package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	tc := "1010101010"
	for i := 1; i < 100000; i++ {
		tc = tc + "1010101010"
	}
	fmt.Println(len(tc))

	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "0001", t: "101"}, want: 1},
		{name: "2", args: args{s: "0101010", t: "1010101"}, want: 7},
		{name: "3", args: args{s: "10101000010011011110", t: "0010011111"}, want: 1},
		{name: "4", args: args{s: "11111111100111111111", t: "100"}, want: 0},
		{name: "5", args: args{s: "0000000000000000000", t: "01111111111111"}, want: 13},
		{name: "6", args: args{s: tc, t: tc}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
