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
		{name: "1", args: args{r: strings.NewReader(`3 5
		3 1
		4 2
		2 3`)}, want: 15},
		{name: "2", args: args{r: strings.NewReader(`4 100
		6 2
		1 5
		3 9
		8 7`)}, want: 100},
		{name: "3", args: args{r: strings.NewReader(`10 3141
		314944731 649
		140276783 228
		578012421 809
		878510647 519
		925326537 943
		337666726 611
		879137070 306
		87808915 39
		756059990 244
		228622672 291`)}, want: 2357689932073},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
