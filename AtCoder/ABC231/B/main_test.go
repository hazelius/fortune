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
		{name: "1", args: args{r: strings.NewReader(`5
		snuke
		snuke
		takahashi
		takahashi
		takahashi`)}, want: "takahashi"},
		{name: "2", args: args{r: strings.NewReader(`5
		takahashi
		takahashi
		aoki
		takahashi
		snuke`)}, want: "takahashi"},
		{name: "3", args: args{r: strings.NewReader(`1
		a`)}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
