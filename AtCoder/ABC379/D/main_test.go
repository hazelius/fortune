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
		{name: "1", args: args{stdin: strings.NewReader(`6
1
2 15
1
3 10
2 20
3 20`)}, wantOut: `1
1
`},
		{name: "2", args: args{stdin: strings.NewReader(`15
1
1
2 226069413
3 1
1
1
2 214168203
1
3 214168203
1
1
1
2 314506461
2 245642315
3 1`)}, wantOut: `2
2
4
`},
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
