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
		{name: "1", args: args{r: strings.NewReader("38 20 17 168 3")}, want: 168},
		{name: "2", args: args{r: strings.NewReader("1 0 1 3 2")}, want: 1},
		{name: "3", args: args{r: strings.NewReader("100 10 100 180 1")}, want: 90},
		{name: "4", args: args{r: strings.NewReader("50 19 20 150 2")}, want: 148},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
