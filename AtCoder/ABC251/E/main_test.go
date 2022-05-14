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
		{name: "1", args: args{r: strings.NewReader(`5
		2 5 3 2 5`)}, want: 7},
		{name: "2", args: args{r: strings.NewReader(`20
		29 27 79 27 30 4 93 89 44 88 70 75 96 3 78 39 97 12 53 62`)}, want: 426},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
