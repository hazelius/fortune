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
		{name: "1", args: args{stdin: strings.NewReader(`7
1 5
1 7
3 2
1 3
1 4
2
3 3`)}, wantOut: `5
10
`},
		{name: "2", args: args{stdin: strings.NewReader(`3
1 1
2
1 3`)}, wantOut: ``},
		{name: "3", args: args{stdin: strings.NewReader(`10
1 15
1 10
1 5
2
1 5
1 10
1 15
2
3 4
3 2`)}, wantOut: `20
5
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
