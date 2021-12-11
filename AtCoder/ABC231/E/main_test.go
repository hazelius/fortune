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
		{name: "1", args: args{r: strings.NewReader(`3 87
		1 10 100`)}, want: 5},
		{name: "2", args: args{r: strings.NewReader(`2 49
		1 7`)}, want: 7},
		{name: "3", args: args{r: strings.NewReader(`10 123456789012345678
		1 100 10000 1000000 100000000 10000000000 1000000000000 100000000000000 10000000000000000 1000000000000000000`)}, want: 233},
		{name: "4", args: args{r: strings.NewReader(`3 60
		1 7 56`)}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
