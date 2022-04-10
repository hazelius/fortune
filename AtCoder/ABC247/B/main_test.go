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
		{name: "1", args: args{r: strings.NewReader(`3
		tanaka taro
		tanaka jiro
		suzuki hanako`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`3
		aaa bbb
		xxx aaa
		bbb yyy`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`2
		tanaka taro
		tanaka taro`)}, want: "No"},
		{name: "4", args: args{r: strings.NewReader(`3
		takahashi chokudai
		aoki kensho
		snu ke`)}, want: "Yes"},
		{name: "5", args: args{r: strings.NewReader(`3
		aaa aaa
		bbb bbb
		ccc ccc`)}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
