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
		{name: "1", args: args{stdin: strings.NewReader(`ch@ku@ai
		choku@@i`)}, wantOut: `Yes`},
		{name: "2", args: args{stdin: strings.NewReader(`ch@kud@i
		akidu@ho`)}, wantOut: `Yes`},
		{name: "3", args: args{stdin: strings.NewReader(`aoki
		@ok@`)}, wantOut: `No`},
		{name: "4", args: args{stdin: strings.NewReader(`aa
		bb`)}, wantOut: `No`},
		{name: "5", args: args{stdin: strings.NewReader(`czaqca@a
		czaqca@w`)}, wantOut: `No`},
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
