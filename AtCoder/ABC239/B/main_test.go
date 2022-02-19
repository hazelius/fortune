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
		{name: "1", args: args{r: strings.NewReader("47")}, want: 4},
		{name: "2", args: args{r: strings.NewReader("-24")}, want: -3},
		{name: "3", args: args{r: strings.NewReader("50")}, want: 5},
		{name: "4", args: args{r: strings.NewReader("-30")}, want: -3},
		{name: "5", args: args{r: strings.NewReader("987654321987654321")}, want: 98765432198765432},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
