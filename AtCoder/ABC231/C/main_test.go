package main

import (
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{r: strings.NewReader(`3 1
		100 160 130
		120`)}, want: "2"},
		{name: "2", args: args{r: strings.NewReader(`5 5
		1 2 3 4 5
		6
		5
		4
		3
		2`)}, want: "0 1 2 3 4"},
		{name: "3", args: args{r: strings.NewReader(`5 5
		804289384 846930887 681692778 714636916 957747794
		424238336
		719885387
		649760493
		596516650
		189641422`)}, want: "5 3 5 5 5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
