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
		{name: "1", args: args{r: strings.NewReader(`7 3
		1 2 1 2 3 3 1`)}, want: 3},
		{name: "2", args: args{r: strings.NewReader(`5 5
		4 4 4 4 4`)}, want: 1},
		{name: "3", args: args{r: strings.NewReader(`10 6
		304621362 506696497 304621362 506696497 834022578 304621362 414720753 304621362 304621362 414720753`)}, want: 4},
		{name: "4", args: args{r: strings.NewReader(`10 3
		2 2 3 3 4 4 5 5 6 6`)}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
