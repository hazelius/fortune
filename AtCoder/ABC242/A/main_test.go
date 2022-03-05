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
		{name: "1", args: args{r: strings.NewReader(`30 500 20 103`)}, want: 0.0425531914893617},
		{name: "2", args: args{r: strings.NewReader(`50 500 100 1`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader(`1 2 1 1000`)}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
