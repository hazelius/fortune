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
		{name: "1", args: args{stdin: strings.NewReader(`11
		ADD 3
		SAVE 1
		ADD 4
		SAVE 2
		LOAD 1
		DELETE
		DELETE
		LOAD 2
		SAVE 1
		LOAD 3
		LOAD 1`)}, wantOut: `3 3 4 4 3 -1 -1 4 4 -1 4 `},
		{name: "2", args: args{stdin: strings.NewReader(`21
		ADD 4
		ADD 3
		DELETE
		ADD 10
		LOAD 7
		SAVE 5
		SAVE 5
		ADD 4
		ADD 4
		ADD 5
		SAVE 5
		ADD 2
		DELETE
		ADD 1
		SAVE 5
		ADD 7
		ADD 8
		DELETE
		ADD 4
		DELETE
		LOAD 5`)}, wantOut: `4 3 4 10 -1 -1 -1 4 4 5 5 2 5 1 1 7 8 7 4 7 1 `},
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
