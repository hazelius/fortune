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
		{name: "1", args: args{r: strings.NewReader("1 0")}, want: 4},
		{name: "2", args: args{r: strings.NewReader("2 5")}, want: 10},
		{name: "3", args: args{r: strings.NewReader("10 10")}, want: 213},
		{name: "4", args: args{r: strings.NewReader("30 100")}, want: 2471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
