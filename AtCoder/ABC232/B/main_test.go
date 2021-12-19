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
		{name: "1", args: args{r: strings.NewReader(`abc
		ijk`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`z
		a`)}, want: "Yes"},
		{name: "3", args: args{r: strings.NewReader(`ppq
		qqp`)}, want: "No"},
		{name: "4", args: args{r: strings.NewReader(`atcoder
		atcoder`)}, want: "Yes"},
		{name: "5", args: args{r: strings.NewReader(`atcoder
		btcoder`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
