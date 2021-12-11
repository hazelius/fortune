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
		want float64
	}{
		{name: "1", args: args{r: strings.NewReader("1000")}, want: 10},
		{name: "2", args: args{r: strings.NewReader("50")}, want: 0.5},
		{name: "3", args: args{r: strings.NewReader("3141")}, want: 31.41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
