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
		sato hanako
		tanaka taro`)}, want: "Yes"},
		{name: "2", args: args{r: strings.NewReader(`3
		saito ichiro
		saito jiro
		saito saburo`)}, want: "No"},
		{name: "3", args: args{r: strings.NewReader(`4
		sypdgidop bkseq
		bajsqz hh
		ozjekw mcybmtt
		qfeysvw dbo`)}, want: "No"},
		{name: "4", args: args{r: strings.NewReader(`2
		tanaka isamu
		tanak aisamu`)}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
