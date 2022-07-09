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
		{name: "1", args: args{r: strings.NewReader("2 2 180")}, want: "-2.0000000000000004 -1.9999999999999998"},
		{name: "2", args: args{r: strings.NewReader("5 0 120")}, want: "-2.499999999999999 4.3301270189221945"},
		{name: "3", args: args{r: strings.NewReader("0 0 11")}, want: "0 0"},
		{name: "4", args: args{r: strings.NewReader("15 5 360")}, want: "15.000000000000002 4.9999999999999964"},
		{name: "5", args: args{r: strings.NewReader("-505 191 278")}, want: "118.85878514480686 526.6674369978655"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
