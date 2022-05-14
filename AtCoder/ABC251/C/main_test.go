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
		aaa 10
		bbb 20
		aaa 30`)}, want: 2},
		{name: "2", args: args{r: strings.NewReader(`5
		aaa 9
		bbb 10
		ccc 10
		ddd 10
		bbb 11`)}, want: 2},
		{name: "3", args: args{r: strings.NewReader(`10
		bb 3
		ba 1
		aa 4
		bb 1
		ba 5
		aa 9
		aa 2
		ab 6
		bb 5
		ab 3`)}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
