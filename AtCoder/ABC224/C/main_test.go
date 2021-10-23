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
		0 1
		1 3
		1 1
		-1 -1`)}, want: 3},
		{name: "2", args: args{r: strings.NewReader(`20
		224 433
		987654321 987654321
		2 0
		6 4
		314159265 358979323
		0 0
		-123456789 123456789
		-1000000000 1000000000
		124 233
		9 -6
		-4 0
		9 5
		-7 3
		333333333 -333333333
		-9 -1
		7 -10
		-1 5
		324 633
		1000000000 -1000000000
		20 0`)}, want: 1124},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
