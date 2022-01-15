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
		{name: "1", args: args{r: strings.NewReader("3 72")}, want: 4},
		{name: "2", args: args{r: strings.NewReader("2 5")}, want: -1},
		{name: "3", args: args{r: strings.NewReader("2 611")}, want: 12},
		{name: "4", args: args{r: strings.NewReader("2 767090")}, want: 111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
