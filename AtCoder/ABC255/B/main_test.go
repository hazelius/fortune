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
		want float64
	}{
		{name: "1", args: args{r: strings.NewReader(`4 2
		2 3
		0 0
		0 1
		1 2
		2 0`)}, want: 2.23606797749978969},
		{name: "2", args: args{r: strings.NewReader(`2 1
		2
		-100000 -100000
		100000 100000`)}, want: 282842.71247461904},
		{name: "3", args: args{r: strings.NewReader(`8 3
		2 6 8
		-17683 17993
		93038 47074
		58079 -57520
		-41515 -89802
		-72739 68805
		24324 -73073
		71049 72103
		47863 19268`)}, want: 130379.280458974768},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
