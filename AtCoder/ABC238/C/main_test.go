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
		{name: "1", args: args{r: strings.NewReader("16")}, want: 73},
		{name: "2", args: args{r: strings.NewReader("238")}, want: 13870},
		{name: "3", args: args{r: strings.NewReader("999999999999999999")}, want: 762062362},
		{name: "4", args: args{r: strings.NewReader("10")}, want: 46},
		{name: "5", args: args{r: strings.NewReader("100")}, want: 4141},
		{name: "6", args: args{r: strings.NewReader("1000")}, want: 409591},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
