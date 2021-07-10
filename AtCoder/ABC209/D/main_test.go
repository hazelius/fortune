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
		{name: "1", args: args{r: strings.NewReader(`4 1
		1 2
		2 3
		2 4
		1 2`)}, want: "Road"},
		{name: "2", args: args{r: strings.NewReader(`5 2
		1 2
		2 3
		3 4
		4 5
		1 3
		1 5`)}, want: "Town\nTown"},
		{name: "3", args: args{r: strings.NewReader(`9 9
		2 3
		5 6
		4 8
		8 9
		4 5
		3 4
		1 9
		3 7
		7 9
		2 5
		2 6
		4 6
		2 4
		5 8
		7 8
		3 6
		5 6`)}, want: "Town\nRoad\nTown\nTown\nTown\nTown\nRoad\nRoad\nRoad"},
		{name: "4", args: args{r: strings.NewReader(`9 1
		2 3
		5 6
		4 8
		8 9
		4 5
		3 4
		1 9
		3 7
		2 4`)}, want: "Town"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
