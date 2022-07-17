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
		{name: "1", args: args{r: strings.NewReader(`6 1 0 2
		80 60 80 60 70 70
		40 20 50 90 90 80`)}, want: "1 4 5"},
		{name: "2", args: args{r: strings.NewReader(`5 2 1 2
		0 100 0 100 0
		0 0 100 100 0`)}, want: "1 2 3 4 5"},
		{name: "3", args: args{r: strings.NewReader(`15 4 3 2
		30 65 20 95 100 45 70 85 20 35 95 50 40 15 85
		0 25 45 35 65 70 80 90 40 55 20 20 45 75 100`)}, want: "2 4 5 6 7 8 11 14 15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
