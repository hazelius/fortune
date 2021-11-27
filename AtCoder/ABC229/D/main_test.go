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
		{name: "1", args: args{r: strings.NewReader(`XX...X.X.X.
		2`)}, want: 5},
		{name: "2", args: args{r: strings.NewReader(`XXXX
		200000`)}, want: 4},
		{name: "3", args: args{r: strings.NewReader(`XX...X..X.XXX.XXX
		3`)}, want: 10},
		{name: "4", args: args{r: strings.NewReader(`.....
		4`)}, want: 4},
		{name: "5", args: args{r: strings.NewReader(`...........X
		2`)}, want: 3},
		{name: "6", args: args{r: strings.NewReader(`X..........
		2`)}, want: 3},
		{name: "7", args: args{r: strings.NewReader(`XX.........X.XX.....
		5`)}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
