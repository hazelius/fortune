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
		{name: "1", args: args{r: strings.NewReader(`5 3
		6 11 2 5 5
		5
		20
		0`)}, want: "10 71 29"},
		{name: "2", args: args{r: strings.NewReader(`10 5
		1000000000 314159265 271828182 141421356 161803398 0 777777777 255255255 536870912 998244353
		555555555
		321654987
		1000000000
		789456123
		0`)}, want: "3316905982 2811735560 5542639502 4275864946 4457360498"},
		{name: "3", args: args{r: strings.NewReader(`5 12
		6 11 2 5 5
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		11
		12`)}, want: "24 19 16 13 10 11 14 17 20 23 26 31"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
