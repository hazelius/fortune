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
		{name: "1", args: args{r: strings.NewReader(`2
		1 8
		4 2`)}, want: "Yes No"},
		{name: "2", args: args{r: strings.NewReader(`4
		201408139683277485 381410962404666524
		360288799186493714 788806911317182736
		18999951915747344 451273909320288229
		962424162689761932 1097438793187620758`)}, want: "No Yes Yes No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.r); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
