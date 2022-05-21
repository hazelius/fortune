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
		{name: "1", args: args{r: strings.NewReader(`3
		1937458062
		8124690357
		2385760149`)}, want: 6},
		{name: "2", args: args{r: strings.NewReader(`5
		0123456789
		0123456789
		0123456789
		0123456789
		0123456789`)}, want: 40},
		{name: "3", args: args{r: strings.NewReader(`3
		0123456789
		9012345678
		2193456780`)}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
