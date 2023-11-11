package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		stdin io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{name: "1", args: args{stdin: strings.NewReader(`6 200
		100 675 201 200 199 328`)}, wantOut: `499`},
		{name: "2", args: args{stdin: strings.NewReader(`8 675
		675 675 675 675 675 675 675 675`)}, wantOut: `5400`},
		{name: "3", args: args{stdin: strings.NewReader(`8 674
		675 675 675 675 675 675 675 675`)}, wantOut: `0`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.stdin, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
