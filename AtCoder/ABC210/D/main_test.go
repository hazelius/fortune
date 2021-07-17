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
		{name: "1", args: args{r: strings.NewReader(`3 4 2
			1 7 7 9
			9 6 3 7
			7 8 6 4`)}, want: 10},
		{name: "2", args: args{r: strings.NewReader(`3 3 1000000000
			1000000 1000000 1
			1000000 1000000 1000000
			1 1000000 1000000`)}, want: 1001000001},
		{name: "3", args: args{r: strings.NewReader(`3 4 2
			100 107 107 1
			1 116 113 117
			111 118 116 114`)}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
