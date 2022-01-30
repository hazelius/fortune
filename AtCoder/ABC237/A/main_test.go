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
		{name: "1", args: args{r: strings.NewReader("10")}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader("-9876543210")}, want: "No"},
		{name: "3", args: args{r: strings.NewReader("483597848400000")}, want: "No"},
		{name: "4", args: args{r: strings.NewReader("2147483648")}, want: "No"},
		{name: "5", args: args{r: strings.NewReader("2147483647")}, want: "Yes"},
		{name: "6", args: args{r: strings.NewReader("-2147483647")}, want: "Yes"},
		{name: "7", args: args{r: strings.NewReader("-2147483648")}, want: "Yes"},
		{name: "8", args: args{r: strings.NewReader("-2147483649")}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
