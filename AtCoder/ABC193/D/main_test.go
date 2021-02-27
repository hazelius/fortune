package main

import (
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		k string
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "1", args: args{k: "2", s: "1144#", t: "2233#"}, want: 0.4444444444444444},
		{name: "2", args: args{k: "2", s: "9988#", t: "1122#"}, want: 1},
		{name: "3", args: args{k: "6", s: "1122#", t: "2228#"}, want: 0.001932367149758454},
		{name: "3", args: args{k: "100000", s: "3226#", t: "3597#"}, want: 0.6296297942426157},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.k, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{a: 2, b: 0}, want: 2},
		{name: "2", args: args{a: 2, b: 1}, want: 20},
		{name: "3", args: args{a: 2, b: 2}, want: 200},
		{name: "4", args: args{a: 3, b: 1}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := point(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("point() = %v, want %v", got, tt.want)
			}
		})
	}
}
