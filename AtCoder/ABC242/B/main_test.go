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
		{name: "1", args: args{r: strings.NewReader(`aba`)}, want: "aab"},
		{name: "2", args: args{r: strings.NewReader(`zzzz`)}, want: "zzzz"},
		{name: "3", args: args{r: strings.NewReader(`bdiabbkhtpzk`)}, want: "abbbdhikkptz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
