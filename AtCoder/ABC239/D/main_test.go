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
		{name: "1", args: args{r: strings.NewReader("2 3 3 4")}, want: "Aoki"},
		{name: "2", args: args{r: strings.NewReader("1 100 50 60")}, want: "Takahashi"},
		{name: "3", args: args{r: strings.NewReader("3 14 1 5")}, want: "Aoki"},
		{name: "4", args: args{r: strings.NewReader("98 99 98 99")}, want: "Aoki"},
		{name: "4", args: args{r: strings.NewReader("1 100 1 100")}, want: "Aoki"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
