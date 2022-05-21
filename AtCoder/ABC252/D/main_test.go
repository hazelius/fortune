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
		{name: "1", args: args{r: strings.NewReader(`4
		3 1 4 1`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`10
		99999 99998 99997 99996 99995 99994 99993 99992 99991 99990`)}, want: 120},
		{name: "3", args: args{r: strings.NewReader(`15
		3 1 4 1 5 9 2 6 5 3 5 8 9 7 9`)}, want: 355},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
