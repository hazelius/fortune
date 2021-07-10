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
		want int
	}{
		{name: "1", args: args{r: strings.NewReader(`2
		1 3`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`4
		3 3 4 4`)}, want: 12},
		{name: "3", args: args{r: strings.NewReader(`2
		1 1`)}, want: 0},
		{name: "4", args: args{r: strings.NewReader(`10
		999999917 999999914 999999923 999999985 999999907 999999965 999999914 999999908 999999951 999999979`)}, want: 405924645},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
