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
		{name: "1", args: args{r: strings.NewReader(`mari
		to
		zzo
		1321`)}, want: "marizzotomari"},
		{name: "2", args: args{r: strings.NewReader(`abra
		cad
		abra
		123`)}, want: "abracadabra"},
		{name: "3", args: args{r: strings.NewReader(`a
		b
		c
		1`)}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
